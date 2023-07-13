package camera

import (
	"context"
	"errors"
	"fmt"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
	"github.com/golang/geo/r3"
	"github.com/viamrobotics/gostream"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage/transform"
	"go.viam.com/rdk/spatialmath"
	"go.viam.com/rdk/utils"
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
	// TODO: 99% confidence this conversion is wrong
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.msg == nil {
		return nil, errors.New("lidar is not ready")
	}

	pc := pointcloud.New()
	numScans := len(l.msg.Ranges)
	angleIncrement := l.msg.AngleIncrement

	for i := 0; i < numScans; i++ {
		lRange := l.msg.Ranges[i]
		// intesity := l.msg.Intensities[i]

		if lRange < l.msg.RangeMin || lRange > l.msg.RangeMax {
			// skip this range
			continue
		}

		err := pc.Set(pointFrom(float64(angleIncrement), utils.DegToRad(0), float64(lRange), 255))
		if err != nil {
			return nil, err
		}

		if pc.Size() == 0 {
			return nil, nil
		}
		return pc, nil
	}

	return nil, fmt.Errorf("not implemented")
}

func pointFrom(yaw, pitch, distance float64, reflectivity uint8) (r3.Vector, pointcloud.Data) {
	ea := spatialmath.NewEulerAngles()
	ea.Yaw = yaw
	ea.Pitch = pitch

	pose1 := spatialmath.NewPose(r3.Vector{X: 0, Y: 0, Z: 0}, ea)
	pose2 := spatialmath.NewPoseFromPoint(r3.Vector{X: distance, Y: 0, Z: 0})
	p := spatialmath.Compose(pose1, pose2).Point()

	// Rotate the point 180 degrees on the y axis. Since lidar data is always 2D, we don't worry
	// about the Z value.
	p.X = -p.X

	pos := pointcloud.NewVector(p.X*1000, p.Y*1000, p.Z*1000)
	d := pointcloud.NewBasicData()
	d.SetIntensity(uint16(reflectivity) * 255)

	return pos, d
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
