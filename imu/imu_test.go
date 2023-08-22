package imu

import (
	"fmt"
	"testing"

	"go.viam.com/test"
)

func TestImuConvert(t *testing.T) {
	msgs, err := loadMessages("data/imu_scan.bag")
	test.That(t, err, test.ShouldBeNil)

	for idx, mm := range msgs {
		fmt.Printf("%d %v\n", idx, mm)
	}

}
