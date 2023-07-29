package camera

import (
	"fmt"
	"os"
	"testing"

	"go.viam.com/rdk/pointcloud"
	"go.viam.com/test"
)

func TestLidarConvert(t *testing.T) {
	msgs, err := loadMessages("data/lidar_laser_scan.bag")
	test.That(t, err, test.ShouldBeNil)

	for idx, mm := range msgs {
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

