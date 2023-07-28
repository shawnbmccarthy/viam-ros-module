package camera

import (
	"testing"

	"go.viam.com/rdk/pointcloud"
	"go.viam.com/test"

	"github.com/shawnbmccarthy/viam-ros-module/utils"
)

func TestLidarConvert(t *testing.T) {
	pc, err := convertMsg(&utils.SampleLaserMsg)
	test.That(t, err, test.ShouldBeNil)


	err = pointcloud.WriteToLASFile(pc, "/tmp/foo.las")
}
