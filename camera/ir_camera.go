package camera

import (
	"context"
	"errors"
	"fmt"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
	"github.com/viamrobotics/gostream"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage/transform"
	"strings"
	"sync"
)

var IrCameraModel = resource.NewModel("viamlabs", "ros", "lidar")

type IrCamera struct {
	resource.Named

	mu         sync.Mutex
	nodeName   string
	primaryUri string
	topic      string
	node       *goroslib.Node
	subscriber *goroslib.Subscriber
	msg        *sensor_msgs.Image
	logger     golog.Logger
}

func init() {
	resource.RegisterComponent(
		camera.API,
		ROSLidarModel,
		resource.Registration[camera.Camera, *IrCameraCfg]{
			Constructor: NewIrCamera,
		},
	)
}

func NewIrCamera(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (camera.Camera, error) {
	c := &IrCamera{
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
	}

	if err := c.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return c, nil
}

// Reconfigure clean this up
func (c *IrCamera) Reconfigure(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
) error {
	var err error
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nodeName = conf.Attributes.String("node_name")
	c.primaryUri = conf.Attributes.String("primary_uri")
	c.topic = conf.Attributes.String("topic")

	if len(strings.TrimSpace(c.primaryUri)) == 0 {
		return errors.New("ROS primary uri must be set to hostname:port")
	}

	if len(strings.TrimSpace(c.topic)) == 0 {
		return errors.New("ROS topic must be set to valid imu topic")
	}

	if len(strings.TrimSpace(c.nodeName)) == 0 {
		c.nodeName = "viam_lidar_node"
	}

	if c.subscriber != nil {
		if c.subscriber.Close() != nil {
			c.logger.Warn("failed to close subscriber")
		}
	}

	if c.node != nil {
		if c.node.Close() != nil {
			c.logger.Warn("failed to close node")
		}
	}

	c.node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          c.nodeName,
		MasterAddress: c.primaryUri,
	})

	// publisher for twist messages
	c.subscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     c.node,
		Topic:    c.topic,
		Callback: c.processMessage,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *IrCamera) processMessage(msg *sensor_msgs.Image) {
	l.msg = msg
}

func (l *IrCamera) Projector(_ context.Context) (transform.Projector, error) {
	return nil, fmt.Errorf("not implemented")
}

func (l *IrCamera) Stream(_ context.Context, _ ...gostream.ErrorHandler) (gostream.VideoStream, error) {
	return nil, fmt.Errorf("not implemented")
}

func (l *IrCamera) NextPointCloud(_ context.Context) (pointcloud.PointCloud, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *IrCamera) Properties(_ context.Context) (camera.Properties, error) {
	return camera.Properties{}, errors.New("not implemented")
}

func (c *IrCamera) Close(_ context.Context) error {
	if c.subscriber.Close() != nil {
		c.logger.Warn("problem closing subscriber")
	}

	if c.node.Close() != nil {
		c.logger.Warn("problem closing node")
	}
	return nil
}
