package utils

import (
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/edaniels/golog"
)

const LASER_SCAN_INVALID = -1.0
const LASER_SCAN_MIN_RANGE = -2.0
const LASER_SCAN_MAX_RANGE = -3.0

type ChannelOption struct {
	NONE      int32 // enable no channels
	INTENSITY int32 // enable intensity channel
	INDEX     int32 // enable index channel
	DISTANCE  int32 // enable distance channel
	TIMESTAMP int32 // enable timestamp channel
	VIEWPOINT int32 // enable viewpoint channel
	DEFAULT   int32 // default channel (INTENSITY | INDEX)
}

func NewChannelOption() *ChannelOption {
	c := &ChannelOption{
		NONE:      0x00,
		INTENSITY: 0x01,
		INDEX:     0x02,
		DISTANCE:  0x04,
		TIMESTAMP: 0x08,
		VIEWPOINT: 0x10,
		DEFAULT:   0x00,
	}

	c.DEFAULT = c.INTENSITY | c.INDEX
	return c
}

type LaserProjection struct {
	angleMin  float32
	angleMax  float32
	conSinMap [][]float32
	logger    golog.Logger
}

func NewProjectLaser() *LaserProjection {
	lp := &LaserProjection{
		angleMin:  0.0,
		angleMax:  0.0,
		conSinMap: [][]float32{{}},
	}
	return lp
}

func (lp *LaserProjection) ProjectLaser(
	scanIn sensor_msgs.LaserScan,
	rangeCutoff float64,
	channelOpts ChannelOption,
) (sensor_msgs.PointCloud2, error) {
	pc := sensor_msgs.PointCloud2{}
	N := len(scanIn.Ranges) // need to sort

	// TODO: is this clean?
	if lp.angleMin != scanIn.AngleMin || lp.angleMax != scanIn.AngleMax {
		lp.logger.Debug("no precomputed map given, computing one")

		lp
	}
	return pc, nil
}
