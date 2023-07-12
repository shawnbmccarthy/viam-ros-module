package lidar

import (
	"context"
	"errors"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
	"github.com/shawnbmccarthy/viam-yahboom-transbot-ros/utils"
	"github.com/viamrobotics/gostream"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage/transform"
	"strings"
	"sync"
	"time"
)

var ROSLidarModel = resource.NewModel("viamlabs", "ros", "lidar")

type ROSLidar struct {
	resource.Named

	mu              sync.Mutex
	nodeName        string
	primaryUri      string
	topic           string
	timeRate        time.Duration // ms to publish
	node            *goroslib.Node
	subscriber      *goroslib.Subscriber
	msg             *sensor_msgs.PointCloud2
	logger          golog.Logger
	laserProjection *utils.LaserProjection
}

func init() {
	resource.RegisterComponent(
		camera.API,
		ROSLidarModel,
		resource.Registration[camera.Camera, *ROSLidarConfig]{
			Constructor: NewROSLidar,
		},
	)
}

func NewROSLidar(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (camera.Camera, error) {
	l := &ROSLidar{
		Named:           conf.ResourceName().AsNamed(),
		logger:          logger,
		laserProjection: utils.NewLaserProjection(),
	}

	if err := l.Reconfigure(ctx, deps, conf); err != nil {
		return nil, err
	}

	return l, nil
}

// Reconfigure clean this up
func (l *ROSLidar) Reconfigure(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
) error {
	var err error
	l.mu.Lock()
	defer l.mu.Unlock()
	l.nodeName = conf.Attributes.String("node_name")
	l.primaryUri = conf.Attributes.String("primary_uri")
	l.topic = conf.Attributes.String("topic")

	if len(strings.TrimSpace(l.primaryUri)) == 0 {
		return errors.New("ROS primary uri must be set to hostname:port")
	}

	if len(strings.TrimSpace(l.topic)) == 0 {
		return errors.New("ROS topic must be set to valid imu topic")
	}

	if len(strings.TrimSpace(l.nodeName)) == 0 {
		l.nodeName = "viam_lidar_node"
	}

	if l.subscriber != nil {
		if l.subscriber.Close() != nil {
			l.logger.Warn("failed to close subscriber")
		}
	}

	if l.node != nil {
		if l.node.Close() != nil {
			l.logger.Warn("failed to close node")
		}
	}

	l.node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          l.nodeName,
		MasterAddress: l.primaryUri,
	})

	// publisher for twist messages
	l.subscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     l.node,
		Topic:    l.topic,
		Callback: l.processMessage,
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *ROSLidar) processMessage(msg *sensor_msgs.LaserScan) {
	msg1 := l.laserProjection.ProjectLaser(msg, -1.0, utils.DEFAULT)
	l.msg = &msg1
}

func (l *ROSLidar) Projector(ctx context.Context) (transform.Projector, error) {
	//TODO implement me
	panic("implement me")
}

func (l *ROSLidar) Stream(ctx context.Context, errHandlers ...gostream.ErrorHandler) (gostream.VideoStream, error) {
	//TODO implement me
	panic("implement me")
}

func (l *ROSLidar) NextPointCloud(ctx context.Context) (pointcloud.PointCloud, error) {
	//TODO implement me
	panic("implement me")
}

func (l *ROSLidar) Properties(ctx context.Context) (camera.Properties, error) {
	//TODO implement me
	panic("implement me")
}

func (l *ROSLidar) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
