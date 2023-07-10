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

var TrackedBaseModel = resource.NewModel("viamlabs", "ros", "trackedbase")

type TrackedBase struct {
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
}

func init() {
	resource.RegisterComponent(
		viambase.API,
		TrackedBaseModel,
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
	t := &TrackedBase{
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
				t.logger.Infof("writing message %v", t.twistMsg)
				t.publisher.Write(t.twistMsg)
			}
		}
	}()

	return t, nil
}

// Reconfigure clean this up
func (t *TrackedBase) Reconfigure(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
) error {
	var err error
	t.mu.Lock()
	defer t.mu.Unlock()
	t.nodeName = conf.Attributes.String("node_name")
	t.primaryUri = conf.Attributes.String("primary_uri")
	t.topic = conf.Attributes.String("topic")

	timeMs := conf.Attributes.Int("time_rate_ms", 250)

	// TODO: create a 250 millisecond publisher?
	t.timeRate = time.Duration(timeMs) * time.Millisecond

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
		if t.publisher.Close() != nil {
			t.logger.Warn("failed to close subscriber")
		}
	}

	if t.node != nil {
		if t.node.Close() != nil {
			t.logger.Warn("failed to close node")
		}
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

func (t *TrackedBase) MoveStraight(
	ctx context.Context,
	distanceMm int,
	mmPerSec float64,
	extra map[string]interface{},
) error {
	t.logger.Warn("MoveStraight Not Implemented")
	return nil
}

func (t *TrackedBase) Spin(
	ctx context.Context,
	angleDeg,
	degsPerSec float64,
	extra map[string]interface{},
) error {
	t.logger.Warn("Spin Not Implemented")
	return nil
}

func (t *TrackedBase) SetPower(
	ctx context.Context,
	linear r3.Vector,
	angular r3.Vector,
	extra map[string]interface{},
) error {
	t.twistMsg.Linear = geometry_msgs.Vector3{X: linear.Y, Y: 0.0, Z: 0.0}
	t.twistMsg.Angular = geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: angular.Z}
	return nil
}

func (t *TrackedBase) SetVelocity(
	ctx context.Context,
	linear r3.Vector,
	angular r3.Vector,
	extra map[string]interface{},
) error {
	t.twistMsg.Linear = geometry_msgs.Vector3{X: linear.Y, Y: 0.0, Z: 0.0}
	t.twistMsg.Angular = geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: angular.Z}
	return nil
}

func (t *TrackedBase) Stop(ctx context.Context, extra map[string]interface{}) error {
	t.twistMsg.Linear = geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: 0.0}
	t.twistMsg.Angular = geometry_msgs.Vector3{X: 0.0, Y: 0.0, Z: 0.0}
	return nil
}

func (t *TrackedBase) IsMoving(ctx context.Context) (bool, error) {
	return false, nil
}

func (t *TrackedBase) Close(_ context.Context) error {
	atomic.StoreInt32(&t.closed, 1)
	err := t.publisher.Close()
	err = t.node.Close()

	return err
}

func (t *TrackedBase) Properties(ctx context.Context, extra map[string]interface{}) (viambase.Properties, error) {
	return viambase.Properties{
		TurningRadiusMeters: 0.0,
		WidthMeters:         0.0,
	}, nil
}

func (t *TrackedBase) Geometries(ctx context.Context) ([]spatialmath.Geometry, error) {
	t.logger.Warn("Geometries Not Implemented")
	return nil, nil
}