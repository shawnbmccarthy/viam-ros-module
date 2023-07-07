package transbot_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type Battery struct {
	msg.Package `ros:"transbot_msgs"`
	Header      std_msgs.Header
	Voltage     float32
}
