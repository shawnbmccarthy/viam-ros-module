package camera

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/bluenviron/goroslib/v2"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
	viamcamera "go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/resource"
	"image"
	"image/color"
	"strings"
	"sync"
)

var RosCameraModel = resource.NewModel("viamlabs", "ros", "camera")

type RosImage struct {
	width  int
	height int
	step   int
	data   []byte
}

func (rosImage *RosImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (rosImage *RosImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: rosImage.height, Y: rosImage.width},
	}
}

func (rosImage *RosImage) At(x, y int) color.RGBA {
	bytesPerPixel := rosImage.step / rosImage.width
	pixelOffset := rosImage.width*x + y
	byteOffset := bytesPerPixel * pixelOffset
	
	return color.RGBA{
		R: rosImage.data[byteOffset+2],
		G: rosImage.data[byteOffset+1],
		B: rosImage.data[byteOffset],
		A: 0,
	}
}

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
	//msg        *RosImage
	img        image.Image
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
	if rs.img != nil {
		return rs.img, func() {}, nil
	} else {
		return nil, nil, fmt.Errorf("image is not ready")
	}
}

func (rs *RosMediaSource) Close(_ context.Context) error {
	// gets closed when stream is closed
	return nil
}

func (rs *RosMediaSource) updateImageFromRosMsg(msg *sensor_msgs.Image) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	//var err error

	if msg == nil || len(msg.Data) == 0 {
		rs.logger.Warn("ROS image data invalid")
		return
	}

	var newData bytes.Buffer
	newData.Write(msg.Data)
	ri := RosImage{height: int(msg.Height), width: int(msg.Width), step: int(msg.Step), data: newData.Bytes()}

	newImage := image.NewRGBA(ri.Bounds())
	for x := 0; x<int(msg.Height); x++ {
		for y := 0; y < int(msg.Width); y++ {
			newImage.Set(x,y,ri.At(x,y))
		}
	}
	rs.img = newImage

		/*
		rs.msg = &RosImage{height: int(msg.Height), width: int(msg.Width), step: int(msg.Step), data: msg.Data}
		buffer := bytes.Buffer{}
		pngEncoder := png.Encoder{CompressionLevel: png.BestCompression}
		var err error
		err = pngEncoder.Encode(&buffer, rs.msg)
		rs.logger.Infof("buffer data: %v", buffer.Bytes())
		if err != nil {
			rs.logger.Errorf("image encoder error: %v", err)
		}

		rs.img, err = png.Decode(&buffer)
		if err != nil {
			rs.logger.Errorf("image decoder error: %v", err)
		} else {
			rs.logger.Infof("info image: %v", rs.img)
		}
		*/

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
