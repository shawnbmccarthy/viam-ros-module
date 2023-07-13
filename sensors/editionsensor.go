package sensors

import (
	"context"
	"errors"
	"github.com/bluenviron/goroslib/v2"
	"github.com/edaniels/golog"
	"github.com/shawnbmccarthy/viam-ros-module/pkg/msgs/transbot_msgs"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/resource"
	"strings"
	"sync"
)

var EditionModel = resource.NewModel("viamlabs", "ros", "yahboomedition")

type EditionSensor struct {
	resource.Named

	mu         sync.Mutex
	nodeName   string
	primaryUri string
	topic      string
	node       *goroslib.Node
	subscriber *goroslib.Subscriber
	msg        *transbot_msgs.Edition
	logger     golog.Logger
}

func init() {
	resource.RegisterComponent(
		sensor.API,
		EditionModel,
		resource.Registration[sensor.Sensor, *EditionSensorConfig]{
			Constructor: NewEditionSensor,
		},
	)
}

func NewEditionSensor(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (sensor.Sensor, error) {
	e := &EditionSensor{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	if err := e.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *EditionSensor) Reconfigure(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.nodeName = conf.Attributes.String("node_name")
	e.primaryUri = conf.Attributes.String("primary_uri")
	e.topic = conf.Attributes.String("topic")

	if len(strings.TrimSpace(e.primaryUri)) == 0 {
		return errors.New("ROS primary uri must be set to hostname:port")
	}

	if len(strings.TrimSpace(e.topic)) == 0 {
		return errors.New("ROS topic must be set to valid imu topic")
	}

	if len(strings.TrimSpace(e.nodeName)) == 0 {
		e.nodeName = "viam_batterysensor_node"
	}

	if e.subscriber != nil {
		if e.subscriber.Close() != nil {
			e.logger.Warn("failed to close subscriber")
		}
	}

	if e.node != nil {
		if e.node.Close() != nil {
			e.logger.Warn("failed to close node")
		}
	}

	var err error
	e.node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          e.nodeName,
		MasterAddress: e.primaryUri,
	})

	if err != nil {
		return err
	}

	e.subscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     e.node,
		Topic:    e.topic,
		Callback: e.processMessage,
	})

	if err != nil {
		return err
	}

	return nil
}

func (e *EditionSensor) processMessage(msg *transbot_msgs.Edition) {
	e.msg = msg
}

func (e *EditionSensor) Readings(
	_ context.Context,
	_ map[string]interface{},
) (map[string]interface{}, error) {
	if e.msg == nil {
		return nil, errors.New("edition message not prepared")
	}
	return map[string]interface{}{"voltage": e.msg.Edition}, nil
}

func (e *EditionSensor) Close(_ context.Context) error {
	err := e.subscriber.Close()
	err = e.node.Close()

	return err
}
