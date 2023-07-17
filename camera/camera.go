package camera

import (
	"context"
	"errors"
	"fmt"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
	viamcamera "go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage"
	"image"
	"strings"
	"sync"
)

var RosCameraModel = resource.NewModel("viamlabs", "ros", "camera")

func NewRosMediaSource(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (*RosMediaSource, error) {
	rms := &RosMediaSource{
		Named:  conf.ResourceName().AsNamed(),
		ctx:    ctx,
		logger: logger,
	}

	if err := rms.Reconfigure(ctx, deps, conf); err != nil {
		logger.Error("problem with reconfigure")
		return nil, err
	}
	return rms, nil
}

type RosMediaSource struct {
	resource.Named

	ctx        context.Context
	logger     golog.Logger
	mu         sync.Mutex
	msg        image.Image
	nodeName   string
	primaryUri string
	topic      string
	node       *goroslib.Node
	subscriber *goroslib.Subscriber
}

func (rs *RosMediaSource) Reconfigure(
	_ context.Context,
	_ resource.Dependencies,
	conf resource.Config,
) error {
	var err error
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.nodeName = conf.Attributes.String("node_name")
	rs.primaryUri = conf.Attributes.String("primary_uri")
	rs.topic = conf.Attributes.String("topic")

	if len(strings.TrimSpace(rs.primaryUri)) == 0 {
		return errors.New("ROS primary uri must be set to hostname:port")
	}

	if len(strings.TrimSpace(rs.topic)) == 0 {
		return errors.New("ROS topic must be set to valid camera topic")
	}

	if len(strings.TrimSpace(rs.nodeName)) == 0 {
		return errors.New("ROS node name must be set for camera node")
	}

	if rs.subscriber != nil {
		if rs.subscriber.Close() != nil {
			rs.logger.Warn("failed to close subscriber")
		}
	}

	if rs.node != nil {
		if rs.node.Close() != nil {
			rs.logger.Warn("failed to close node")
		}
	}

	rs.node, err = goroslib.NewNode(goroslib.NodeConf{
		Name:          rs.nodeName,
		MasterAddress: rs.primaryUri,
	})
	if err != nil {
		rs.logger.Errorf("problem creating node: %v", err)
		return err
	}

	// publisher for twist messages
	rs.subscriber, err = goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     rs.node,
		Topic:    rs.topic,
		Callback: rs.updateImageFromRosMsg,
	})
	if err != nil {
		rs.logger.Errorf("problem creating subscriber: %v", err)
		return err
	}

	return nil
}

func (rs *RosMediaSource) Read(_ context.Context) (image.Image, func(), error) {
	return rs.msg, func() {}, nil
}

func (rs *RosMediaSource) Close(_ context.Context) error {
	return nil
}

func (rs *RosMediaSource) updateImageFromRosMsg(msg *sensor_msgs.Image) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	var err error

	if msg != nil || len(msg.Data) > 0 {
		//TODO: really bad loop
		for i := 0; i < len(msg.Data); {
			msg.Data[i], msg.Data[i+2] = msg.Data[i+2], msg.Data[i]
			i = i + 3
		}
		rs.msg, err = rimage.DecodeImage(rs.ctx, msg.Data, "png")
		if err != nil {
			rs.logger.Errorf("Cannot decode immage %v", err)
		}
	} else {
		rs.logger.Warn("ROS image data invalid")
	}
}

type RosMediaSourceConfig struct {
	NodeName   string `json:"node_name"`
	PrimaryUri string `json:"primary_uri"`
	Topic      string `json:"topic"`
}

func (cfg *RosMediaSourceConfig) Validate(path string) ([]string, error) {
	// NodeName will get default value if string is empty
	if cfg.PrimaryUri == "" {
		return nil, fmt.Errorf(`expected "PrimaryUri" attribute for sensor %q`, path)
	}

	if cfg.Topic == "" {
		return nil, fmt.Errorf(`expected "RosTopic" attribute for sensor %q`, path)
	}

	return nil, nil
}

func init() {
	resource.RegisterComponent(
		viamcamera.API,
		RosCameraModel,
		resource.Registration[viamcamera.Camera, *RosMediaSourceConfig]{
			Constructor: NewRosCamera,
		},
	)
}

func NewRosCamera(
	ctx context.Context,
	deps resource.Dependencies,
	conf resource.Config,
	logger golog.Logger,
) (viamcamera.Camera, error) {
	rosVideoSrc, err := NewRosMediaSource(ctx, deps, conf, logger)
	if err != nil {
		logger.Error("problem creating ROS media source")
		return nil, err
	}

	camModel := viamcamera.NewPinholeModelWithBrownConradyDistortion(nil, nil)
	videoSrc, err := viamcamera.NewVideoSourceFromReader(
		ctx,
		rosVideoSrc,
		&camModel,
		viamcamera.ColorStream,
	)
	if err != nil {
		logger.Error("problem created new video source")
		return nil, err
	}
	return viamcamera.FromVideoSource(conf.ResourceName(), videoSrc), nil
}
