package base

import (
	"context"
	"errors"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	viambase "go.viam.com/rdk/components/base"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var RosBaseModel = resource.NewModel("viamlabs", "ros", "base")

type RosBase struct {
	resource.Named

	mu         sync.Mutex
	nodeName   string
	primaryUri string
	topic      string
	timeRate   time.Duration // ms to publish
	node       *goroslib.Node
	publisher  *goroslib.Publisher
	twistMsg   *geometry_msgs.Twist
	logger     golog.Logger
	msgRate    *goroslib.NodeRate
	closed     int32
	moving     bool
}

func init() {
	resource.RegisterComponent(
		viambase.API,
		RosBaseModel,
		resource.Registration[viambase.Base, *TrackedBaseConfig]{
			Constructor: NewTrackedBase,
		},
	)
}

func NewTrackedBase(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (viambase.Base, error) {
	t := &RosBase{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	if err := t.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	// set twist message thread
	go func() {
		for atomic.LoadInt32(&t.closed) == 0 {
			select {
			case <-t.msgRate.SleepChan():
				t.mu.Lock()
				t.publisher.Write(t.twistMsg)
				t.mu.Unlock()
			}
		}
	}()

	return t, nil
}

// Reconfigure clean this up
func (t *RosBase) Reconfigure(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
) error {
	var err error
	t.mu.Lock()
	defer t.mu.Unlock()
	t.nodeName = conf.Attributes.String("node_name")
	t.primaryUri = conf.Attributes.String("primary_uri")
	t.topic = conf.Attributes.String("topic")

	timeMs := conf.Attributes.Int("time_rate_ms", 250)

	t.timeRate = time.Duration(timeMs) * time.Millisecond
	t.twistMsg = &geometry_msgs.Twist{
		Linear:  geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Angular: geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
	}

	if len(strings.TrimSpace(t.primaryUri)) == 0 {
		return errors.New("ROS primary uri must be set to hostname:port")
	}

	if len(strings.TrimSpace(t.topic)) == 0 {
		return errors.New("ROS topic must be set to valid imu topic")
	}

	if len(strings.TrimSpace(t.nodeName)) == 0 {
		t.nodeName = "viam_trackedbase_node"
	}

	if t.publisher != nil {
		t.publisher.Close()
	}

	if t.node != nil {
		t.node.Close()
	}

	t.node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          t.nodeName,
		MasterAddress: t.primaryUri,
	})

	t.msgRate = t.node.TimeRate(t.timeRate)
	if err != nil {
		return err
	}

	// publisher for twist messages
	t.publisher, err = goroslib.NewPublisher(goroslib.PublisherConf{
		Node:  t.node,
		Topic: t.topic,
		Msg:   &geometry_msgs.Twist{},
	})
	if err != nil {
		return err
	}
	atomic.StoreInt32(&t.closed, 0)
	return nil
}

func (t *RosBase) MoveStraight(
	ctx context.Context,
	distanceMm int,
	mmPerSec float64,
	extra map[string]interface{},
) error {
	t.logger.Warn("MoveStraight Not Implemented")
	return nil
}

func (t *RosBase) Spin(
	ctx context.Context,
	angleDeg,
	degsPerSec float64,
	extra map[string]interface{},
) error {
	t.logger.Warn("Spin Not Implemented")
	return nil
}

func (t *RosBase) SetPower(
	_ context.Context,
	linear r3.Vector,
	angular r3.Vector,
	_ map[string]interface{},
) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.twistMsg = &geometry_msgs.Twist{
		Linear:  geometry_msgs.Vector3{X: linear.Y, Y: 0.0, Z: 0.0},
		Angular: geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: angular.Z},
	}
	t.moving = true
	return nil
}

func (t *RosBase) SetVelocity(
	_ context.Context,
	linear r3.Vector,
	angular r3.Vector,
	_ map[string]interface{},
) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.twistMsg = &geometry_msgs.Twist{
		Linear:  geometry_msgs.Vector3{X: linear.Y, Y: 0.0, Z: 0.0},
		Angular: geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: angular.Z},
	}
	t.moving = true
	return nil
}

func (t *RosBase) Stop(_ context.Context, _ map[string]interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.twistMsg = &geometry_msgs.Twist{
		Linear:  geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Angular: geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
	}
	t.moving = false
	return nil
}

func (t *RosBase) IsMoving(_ context.Context) (bool, error) {
	return t.moving, nil
}

func (t *RosBase) Close(_ context.Context) error {
	atomic.StoreInt32(&t.closed, 1)
	t.publisher.Close()
	t.node.Close()
	return nil
}

func (t *RosBase) Properties(
	_ context.Context,
	_ map[string]interface{},
) (viambase.Properties, error) {
	return viambase.Properties{
		TurningRadiusMeters: 0.0,
		WidthMeters:         0.0,
	}, nil
}

func (t *RosBase) Geometries(_ context.Context, _ map[string]interface{}) ([]spatialmath.Geometry, error) {
	t.logger.Warn("Geometries Not Implemented")
	return nil, nil
}
