package imu

import (
	"testing"

	"go.viam.com/test"
)

func TestImuConvert(t *testing.T) {
	msgs, err := loadMessages("data/imu.bag")
	// no errors and there should be messages
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(msgs), test.ShouldBeGreaterThan, 1)
}
