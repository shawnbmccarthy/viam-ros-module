package main

import (
	"fmt"
	"github.com/shawnbmccarthy/viam-yahboom-transbot-ros/utils"
)

func main() {
	co := utils.NewChannelOption()
	fmt.Printf("hello: %b\n", co.NONE)
	fmt.Printf("hello: %b\n", co.INTENSITY)
	fmt.Printf("hello: %b\n", co.INDEX)
	fmt.Printf("hello: %b\n", co.DISTANCE)
	fmt.Printf("hello: %b\n", co.TIMESTAMP)
	fmt.Printf("hello: %b\n", co.VIEWPOINT)
	fmt.Printf("hello: %b\n", co.DEFAULT)
	fmt.Printf("hello: %b\n", 0x00)
}
