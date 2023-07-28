package camera

import (
	"context"
	"errors"
	"fmt"
	"math"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
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

	mu         sync.Mutex
	nodeName   string
	primaryUri string
	topic      string
	timeRate   time.Duration // ms to publish
	node       *goroslib.Node
	subscriber *goroslib.Subscriber
	msg        *sensor_msgs.LaserScan
	pcMsg      pointcloud.PointCloud
	logger     golog.Logger
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
		Named:  conf.ResourceName().AsNamed(),
		logger: logger,
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
	l.msg = msg
}

func (l *ROSLidar) Projector(_ context.Context) (transform.Projector, error) {
	return nil, fmt.Errorf("not implemented")
}

func (l *ROSLidar) Stream(_ context.Context, _ ...gostream.ErrorHandler) (gostream.VideoStream, error) {
	return nil, fmt.Errorf("not implemented")
}

func (l *ROSLidar) NextPointCloud(_ context.Context) (pointcloud.PointCloud, error) {

	l.mu.Lock()
	msg := l.msg
	l.mu.Unlock()

	if msg == nil {
		return nil, errors.New("lidar is not ready")
	}

	return convertMsg(msg)
}

func convertMsg(msg *sensor_msgs.LaserScan) (pointcloud.PointCloud, error) {

	pc := pointcloud.New()

	for i, r := range msg.Ranges {
		if r < msg.RangeMin || r > msg.RangeMax {
			// TODO(erh): is this right? needed?
			continue
		}
		
		p := r3.Vector{}
		ang := msg.AngleMin + (float32(i)*msg.AngleIncrement)
		p.X = math.Sin(float64(ang)) * float64(r)
		p.Y = math.Cos(float64(ang)) * float64(r)
		
		d := pointcloud.NewBasicData()
		d.SetIntensity(uint16(msg.Intensities[i]))
		
		err := pc.Set(p, d)
		if err != nil {
			return nil, err
		}
	}

	return pc, nil
}

func (l *ROSLidar) Properties(_ context.Context) (camera.Properties, error) {
	return camera.Properties{}, errors.New("not implemented")
}

func (l *ROSLidar) Close(_ context.Context) error {
	if l.subscriber.Close() != nil {
		l.logger.Warn("problem closing subscriber")
	}

	if l.node.Close() != nil {
		l.logger.Warn("problem closing node")
	}
	return nil
}
