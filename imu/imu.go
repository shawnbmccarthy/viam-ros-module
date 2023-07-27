package imu

import (
	"context"
	"strings"
	"sync"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"
	"github.com/pkg/errors"

	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var Model = resource.NewModel("viamlabs", "ros", "imu")

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
		if r.subscriber.Close() != nil {
			r.logger.Warn("failed to close subscriber")
		}
	}

	if r.node != nil {
		if r.node.Close() != nil {
			r.logger.Warn("failed to close node")
		}
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
	// GPS - not supported
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

	err := r.subscriber.Close()
	if err != nil {
		return err
	}

	err = r.node.Close()
	if err != nil {
		return err
	}

	return nil
}
