package imu

import (
	"context"
	"go.viam.com/rdk/ros"
	"strings"
	"sync"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"
	"github.com/pkg/errors"

	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var Model = resource.NewModel("viamlabs", "ros", "imu")
var DummyImuModel = resource.NewModel("viamlabs", "ros", "imu-dummy")

type RosImu struct {
	resource.Named

	mu         sync.Mutex
	nodeName   string
	primaryUri string
	topic      string
	node       *goroslib.Node
	subscriber *goroslib.Subscriber
	msg        *sensor_msgs.Imu
	logger     golog.Logger
}

func init() {
	resource.RegisterComponent(
		movementsensor.API,
		Model,
		resource.Registration[movementsensor.MovementSensor, *RosImuConfig]{
			Constructor: NewRosImu,
		},
	)

	resource.RegisterComponent(
		movementsensor.API,
		DummyImuModel,
		resource.Registration[movementsensor.MovementSensor, resource.NoNativeConfig]{
			Constructor: NewRosImuDummy,
		},
	)
}

func NewRosImu(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (movementsensor.MovementSensor, error) {
	r := &RosImu{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	if err := r.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return r, nil
}

func NewRosImuDummy(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (movementsensor.MovementSensor, error) {
	r := &RosImu{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}
	msgs, err := loadMessages(conf.Attributes.String("bag"))
	if err != nil {
		return nil, err
	}

	r.msg = &msgs[0]

	return r, nil
}

func (r *RosImu) Reconfigure(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nodeName = conf.Attributes.String("node_name")
	r.primaryUri = conf.Attributes.String("primary_uri")
	r.topic = conf.Attributes.String("topic")

	if len(strings.TrimSpace(r.primaryUri)) == 0 {
		return errors.New("ROS primary uri must be set to hostname:port")
	}

	if len(strings.TrimSpace(r.topic)) == 0 {
		return errors.New("ROS topic must be set to valid imu topic")
	}

	if len(strings.TrimSpace(r.nodeName)) == 0 {
		r.nodeName = "viam_imu_node"
	}

	if r.subscriber != nil {
		r.subscriber.Close()
	}

	if r.node != nil {
		r.node.Close()
	}

	var err error
	r.node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          r.nodeName,
		MasterAddress: r.primaryUri,
	})
	if err != nil {
		return err
	}

	r.subscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     r.node,
		Topic:    r.topic,
		Callback: r.processMessage,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *RosImu) processMessage(msg *sensor_msgs.Imu) {
	r.msg = msg
}

func (r *RosImu) Position(
	ctx context.Context,
	extra map[string]interface{},
) (*geo.Point, float64, error) {
	// TODO: Implement GPS
	return geo.NewPoint(0, 0), 0, movementsensor.ErrMethodUnimplementedPosition
}

func (r *RosImu) LinearVelocity(
	ctx context.Context,
	extra map[string]interface{},
) (r3.Vector, error) {
	// GPS - not supported
	return r3.Vector{}, movementsensor.ErrMethodUnimplementedLinearVelocity
}

func (r *RosImu) AngularVelocity(
	ctx context.Context,
	extra map[string]interface{},
) (spatialmath.AngularVelocity, error) {
	// IMU
	if r.msg == nil {
		return spatialmath.AngularVelocity{}, errors.New("message unavailable")
	}
	av := spatialmath.AngularVelocity{
		X: r.msg.AngularVelocity.X,
		Y: r.msg.AngularVelocity.Y,
		Z: r.msg.AngularVelocity.Z,
	}
	return av, nil
}

func (r *RosImu) LinearAcceleration(
	ctx context.Context,
	extra map[string]interface{},
) (r3.Vector, error) {
	// IMU
	if r.msg == nil {
		return r3.Vector{}, errors.New("message unavailable")
	}
	la := r3.Vector{
		X: r.msg.LinearAcceleration.X,
		Y: r.msg.LinearAcceleration.Y,
		Z: r.msg.LinearAcceleration.Z,
	}
	return la, nil
}

func (r *RosImu) CompassHeading(
	ctx context.Context,
	extra map[string]interface{},
) (float64, error) {
	return 0, movementsensor.ErrMethodUnimplementedCompassHeading
}

func (r *RosImu) Orientation(
	ctx context.Context,
	extra map[string]interface{},
) (spatialmath.Orientation, error) {
	// IMU
	return nil, errors.New("conversion not supported at this time")
}

func (r *RosImu) Properties(
	ctx context.Context,
	extra map[string]interface{},
) (*movementsensor.Properties, error) {
	return &movementsensor.Properties{
		LinearAccelerationSupported: true,
	}, nil
}

func (r *RosImu) Readings(
	ctx context.Context,
	extra map[string]interface{},
) (map[string]interface{}, error) {
	readings, err := movementsensor.Readings(ctx, r, extra)
	if err != nil {
		return nil, err
	}
	return readings, nil
}

func (r *RosImu) Accuracy(
	ctx context.Context,
	extra map[string]interface{},
) (map[string]float32, error) {
	// GPS - not supported
	return map[string]float32{}, movementsensor.ErrMethodUnimplementedAccuracy
}

func (r *RosImu) Close(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.subscriber != nil {
		r.subscriber.Close()
	}

	if r.node != nil {
		r.node.Close()
	}

	return nil
}

func loadMessages(fn string) ([]sensor_msgs.Imu, error) {
	bag, err := ros.ReadBag(fn)
	if err != nil {
		return nil, err
	}

	err = ros.WriteTopicsJSON(bag, 0, 0, nil)
	if err != nil {
		return nil, err
	}

	all, err := ros.AllMessagesForTopic(bag, "imu_data")
	if err != nil {
		return nil, err
	}

	fixed := []sensor_msgs.Imu{}

	for _, m := range all {
		mm := sensor_msgs.Imu{}
		data := m["data"].(map[string]interface{})

		orientation := data["orientation"].(map[string]interface{})
		orientationCovariance := data["orientation_covariance"].([]interface{})
		angularVel := data["angular_velocity"].(map[string]interface{})
		angularVelCovariance := data["angular_velocity_covariance"].([]interface{})
		linearAccel := data["linear_acceleration"].(map[string]interface{})
		linearAccelCovariance := data["linear_acceleration_covariance"].([]interface{})
		// TODO(SM): covariance array conversions

		mm.Orientation = geometry_msgs.Quaternion{
			X: orientation["x"].(float64),
			Y: orientation["y"].(float64),
			Z: orientation["z"].(float64),
			W: orientation["w"].(float64),
		}
		for idx, oc := range orientationCovariance {
			mm.OrientationCovariance[idx] = oc.(float64)
		}

		mm.AngularVelocity = geometry_msgs.Vector3{
			X: angularVel["x"].(float64),
			Y: angularVel["y"].(float64),
			Z: angularVel["z"].(float64),
		}
		for idx, avc := range angularVelCovariance {
			mm.AngularVelocityCovariance[idx] = avc.(float64)
		}

		mm.LinearAcceleration = geometry_msgs.Vector3{
			X: linearAccel["x"].(float64),
			Y: linearAccel["y"].(float64),
			Z: linearAccel["z"].(float64),
		}
		for idx, lac := range linearAccelCovariance {
			mm.LinearAccelerationCovariance[idx] = lac.(float64)
		}

		fixed = append(fixed, mm)
	}
	return fixed, nil
}
