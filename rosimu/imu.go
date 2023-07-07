package rosimu

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	geo "github.com/kellydunn/golang-geo"

	"go.viam.com/rdk/components/movementsensor"
	"go.viam.com/rdk/resource"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
)

var Model = resource.NewModel("viamlabs", "ros", "imu")

type MyMovementSensor struct {
	resource.Named

	mu            sync.Mutex
	rosname       string
	rosmaster     string
	rostopic      string
	rosnode       Node
	rossubscriber Subscriber
	rosmsg        sensor_msgs.Imu
}

type MyMovementSensorConfig struct {
	RosName   string `json:"rosnodename"`
	RosMaster string `json:"rosmaster"`
	RosTopic  string `json:"rostopic"`
}

func init() {
	resource.RegisterComponent(movementsensor.API, Model, resource.Registration[movementsensor.MovementSensor, *MyMovementSensorConfig]{
		Constructor: newMovementSensor,
	})
}

func newMovementSensor(ctx context.Context, deps resource.Dependencies, conf resource.Config, logger golog.Logger) (movementsensor.MovementSensor, error) {
	m := &MyMovementSensor{
		named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	if err := m.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MyMovementSensor) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.rosname = conf.Attributes.String("rosnodename")
	m.rosmaster = conf.Attributes.String("rosmaster")
	m.rostopic = conf.Attributes.String("rostopic")

	if len(strings.TrimSpace(m.rosmaster)) == 0 {
		return errors.New("rosmaster must be set to hostname:port")
	}

	if len(strings.TrimSpace(m.rostopic)) == 0 {
		return errors.New("rostopic must be set to valid imu topic")
	}

	if len(strings.TrimSpace(m.rosnodename)) == 0 {
		m.rosnodename = "viam_imu_node"
	}

	// does a struct set the value to nil first?
	if m.rossubscriber != nil {
		m.rossubscriber.Close()
	}

	if m.rosnode != nil {
		m.rosnode.Close()
	}

	m.rosnode, err := goroslib.NewNode(goroslib.NodeConf{
		Name: m.rosnodename,
		MasterAddress: m.rosmaster,
	})

	m.rossubscriber, err := goroslib.Subscriber{
		Node: m.rosnode,
		Topic: m.rostopic,
		Callback: m.processMessage,
	})

	if err != nil {
		return err
	}

	return
}

func (m *MyMovementSensor) processMessage(msg *sensor_msgs.Imu) {
	return
}

func (m *MyMovementSensor) Position(Position(ctx context.Context, extra map[string]interface{}) (*geo.Point, float64, error) {
	// GPS - not supported
	return geo.NewPoint(0, 0), 0, movementsensor.ErrMethodUnimplementedPosition
}

func (m *MyMovementSensor) LinearVelocity(ctx context.Context, extra map[string]interface{}) (r3.Vector, error) {
	// GPS - not supported
	return r3.Vector{}, movementsensor.ErrMethodUnimplementedLinearVelocity
}

func (m *MyMovementSensor) AngularVelocity(ctx context.Context, extra map[string]interface{}) (spatialmath.AngularVelocity, error) {
	// IMU
	return errUnimplemented
}

func (m *MyMovementSensor) LinearAcceleration(ctx context.Context, extra map[string]interface{}) (r3.Vector, error) {
	// IMU
	return errUnimplemented
}

func (m *MyMovementSensor) CompassHeading(ctx context.Context, extra map[string]interface{}) (float64, error) {
	// GPS - not supported
	return 0, movementsensor.ErrMethodUnimplementedCompassHeading
}

func (m *MyMovementSenosr) Orientation(ctx context.Context, extra map[string]interface{}) (spatialmath.Orientation, error) {
	// IMU
	return errUnimplemented
}

func (m *MyMovementSensor) Properties(ctx context.Context, extra map[string]interface{}) (*Properties, error) {
	return &movementsensor.Properties{
		LinearAccesslerationSupported: true,
	}, nil
}

func (m *MyMovementSensor) Accuracy(ctx context.Context, extra map[string]interface{}) (map[string]float32, error) {
	// GPS - not supported
	return map[string]float32{}, movementsensor.ErrMethodUnimplementedAccuracy
}
