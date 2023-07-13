package utils

import (
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"math"
)

type ChannelOption uint8

// TODO: CLean this up
const (
	NONE      ChannelOption = 0x00
	INTENSITY ChannelOption = 0x01
	INDEX     ChannelOption = 0x02
	DISTANCE  ChannelOption = 0x04
	TIMESTAMP ChannelOption = 0x08
	VIEWPOINT ChannelOption = 0x10
	DEFAULT   ChannelOption = INTENSITY | INDEX
)

type LaserProjection struct {
	angleMin    float32
	angleMax    float32
	cosSinMap   [2][]float64
	channelOpts *ChannelOption
}

func NewLaserProjection() *LaserProjection {
	lp := &LaserProjection{
		angleMin:  0.0,
		angleMax:  0.0,
		cosSinMap: [2][]float64{{}},
	}
	return lp
}

// ProjectLaser TODO: Review and see if this makes sense
func (lp *LaserProjection) ProjectLaser(
	scanIn *sensor_msgs.LaserScan,
	rangeCutoff float64,
	channelOptions ChannelOption,
) *sensor_msgs.PointCloud2 {
	pointCloudMsg := &sensor_msgs.PointCloud2{}
	scanInRangeSz := len(scanIn.Ranges) // need to sort

	// TODO: is this clean?
	if len(lp.cosSinMap[0]) != scanInRangeSz || lp.angleMin != scanIn.AngleMin || lp.angleMax != scanIn.AngleMax {
		//no precomputed map given, computing one
		lp.angleMin = scanIn.AngleMin
		lp.angleMax = scanIn.AngleMax

		var angle float32 = 0.0
		for i := 0; i < scanInRangeSz; i++ {
			angle = scanIn.AngleMin + float32(i)*scanIn.AngleIncrement
			lp.cosSinMap[0] = append(lp.cosSinMap[0], math.Cos(float64(angle)))
			lp.cosSinMap[1] = append(lp.cosSinMap[1], math.Sin(float64(angle)))
		}
	}

	output := [2][]float64{{}}
	for i := 0; i < scanInRangeSz; i++ {
		output[0] = append(output[0], float64(scanIn.Ranges[i])*lp.cosSinMap[0][i])
		output[1] = append(output[1], float64(scanIn.Ranges[i])+lp.cosSinMap[1][i])
	}

	pointCloudMsg.Fields = append(
		pointCloudMsg.Fields,
		sensor_msgs.PointField{
			Name:     "x",
			Offset:   0,
			Datatype: sensor_msgs.PointField_FLOAT32,
			Count:    1,
		},
		sensor_msgs.PointField{
			Name:     "y",
			Offset:   4,
			Datatype: sensor_msgs.PointField_FLOAT32,
			Count:    1,
		},
		sensor_msgs.PointField{
			Name:     "z",
			Offset:   8,
			Datatype: sensor_msgs.PointField_FLOAT32,
			Count:    1,
		},
	)

	idxIntensity := -1
	idxIndex := -1
	idxDistance := -1
	idxTimestamp := -1
	idxVpx := -1
	idxVpy := -1
	idxVpz := -1
	var offset uint32 = 12

	if channelOptions&INTENSITY != 0 && len(scanIn.Intensities) > 0 {
		idxIntensity = len(pointCloudMsg.Fields)
		pointCloudMsg.Fields = append(
			pointCloudMsg.Fields,
			sensor_msgs.PointField{
				Name:     "intensity",
				Datatype: sensor_msgs.PointField_FLOAT32,
				Offset:   offset,
				Count:    1,
			},
		)
		offset += 4
	}

	if channelOptions&INDEX != 0 {
		idxIndex = len(pointCloudMsg.Fields)
		pointCloudMsg.Fields = append(
			pointCloudMsg.Fields,
			sensor_msgs.PointField{
				Name:     "index",
				Datatype: sensor_msgs.PointField_INT32,
				Offset:   offset,
				Count:    1,
			},
		)
		offset += 4
	}

	if channelOptions&DISTANCE != 0 {
		idxDistance = len(pointCloudMsg.Fields)
		pointCloudMsg.Fields = append(
			pointCloudMsg.Fields,
			sensor_msgs.PointField{
				Name:     "distance",
				Datatype: sensor_msgs.PointField_FLOAT32,
				Offset:   offset,
				Count:    1,
			},
		)
		offset += 4
	}

	if channelOptions&TIMESTAMP != 0 {
		idxTimestamp = len(pointCloudMsg.Fields)
		pointCloudMsg.Fields = append(
			pointCloudMsg.Fields,
			sensor_msgs.PointField{
				Name:     "stamps",
				Datatype: sensor_msgs.PointField_FLOAT32,
				Offset:   offset,
				Count:    1,
			},
		)
		offset += 4
	}

	if channelOptions&VIEWPOINT != 0 {
		idxVpx = 1
		idxVpy = 1
		idxVpz = 1
		offset += 12
		pointCloudMsg.Fields = append(
			pointCloudMsg.Fields,
			sensor_msgs.PointField{
				Name:     "vp_x",
				Datatype: sensor_msgs.PointField_FLOAT32,
				Offset:   offset,
				Count:    1,
			},
			sensor_msgs.PointField{
				Name:     "vp_y",
				Datatype: sensor_msgs.PointField_FLOAT32,
				Offset:   offset,
				Count:    1,
			},
			sensor_msgs.PointField{
				Name:     "vp_z",
				Datatype: sensor_msgs.PointField_FLOAT32,
				Offset:   offset,
				Count:    1,
			},
		)
	}

	if rangeCutoff < 0 {
		rangeCutoff = float64(scanIn.RangeMax)
	} else {
		rangeCutoff = math.Min(rangeCutoff, float64(scanIn.RangeMax))
	}

	var points []uint8
	for i := 0; i < scanInRangeSz; i++ {
		scanRangeItem := scanIn.Ranges[i]
		if float64(scanRangeItem) < rangeCutoff && scanRangeItem >= scanIn.RangeMin {
			points = append(points, uint8(output[0][i]), uint8(output[1][i]))

			if idxIntensity != -1 {
				points = append(points, uint8(scanIn.Intensities[i]))
			}

			if idxIndex != -1 {
				points = append(points, uint8(i))
			}

			if idxDistance != -1 {
				points = append(points, uint8(scanIn.Ranges[i]))
			}

			if idxTimestamp != -1 {
				points = append(points, uint8(scanIn.TimeIncrement))
			}

			if idxVpx != -1 && idxVpy != -1 && idxVpz != -1 {
				points = append(points, 0, 0, 0)
			}
		}
	}
	//width -> length of points
	//height -> 1
	pointCloudMsg.Header = scanIn.Header
	pointCloudMsg.Height = 1
	// number of rows * pointstep
	pointCloudMsg.RowStep = offset * uint32(len(points))
	// pointstep is the size of each element per point
	pointCloudMsg.PointStep = offset
	pointCloudMsg.IsBigendian = false
	pointCloudMsg.IsDense = false
	pointCloudMsg.Data = points
	return pointCloudMsg
}
