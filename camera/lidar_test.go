package camera

import (
	"fmt"
	"os"
	"testing"

	"go.viam.com/rdk/pointcloud"
	"go.viam.com/rdk/ros"
	"go.viam.com/test"
	
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	
)

func TestLidarConvert(t *testing.T) {
	bag, err := ros.ReadBag("data/lidar_laser_scan.bag")
	test.That(t, err, test.ShouldBeNil)

	err = ros.WriteTopicsJSON(bag, 0, 0, nil)
	test.That(t, err, test.ShouldBeNil)

	all, err := ros.AllMessagesForTopic(bag, "scan")
	test.That(t, err, test.ShouldBeNil)

	test.That(t, 147, test.ShouldEqual, len(all))

	for idx, m := range all {
		mm := sensor_msgs.LaserScan{}
		// TODO(erh): there must be a better way to do this

		data := m["data"].(map[string]interface{})

		mm.AngleMin = float32(data["angle_min"].(float64))
		mm.AngleMax = float32(data["angle_max"].(float64))
		mm.AngleIncrement = float32(data["angle_increment"].(float64))
		mm.TimeIncrement = float32(data["time_increment"].(float64))
		mm.ScanTime = float32(data["scan_time"].(float64))
		mm.RangeMin = float32(data["range_min"].(float64))
		mm.RangeMax = float32(data["range_max"].(float64))

		mm.Intensities = fixArray(data["intensities"])
		mm.Ranges = fixArray(data["ranges"])
		
		pc, err := convertMsg(&mm)
		test.That(t, err, test.ShouldBeNil)

		fn := fmt.Sprintf("/tmp/foo%d.pcd", idx)
		f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0755)
		test.That(t, err, test.ShouldBeNil)
		defer f.Close()

		fmt.Println(fn)
	
		err = pointcloud.ToPCD(pc, f, pointcloud.PCDBinary)
		test.That(t, err, test.ShouldBeNil)
	}

}

func fixArray(a interface{}) []float32 {
	b := a.([]interface{})
	c := []float32{}
	for _, n := range b {
		c = append(c, float32(n.(float64)))
	}
	return c
}
