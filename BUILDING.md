# Building ROS Module

If you have a custom ROS robot which you would like to integrate with Viam, you can clone this repo and extend the system 
as needed.

The initial ROS integration only supports:
1. a few yahboom custom messages to demonstrate how we can integrate custom ROS messages from the yahboom overlay.
2. ROS Melodic & ROS Noetic (for Noetic we are still running some tests)

## Prerequisites

1. Install [go](https://go.dev/learn/), this can be on your ROS system or another system (we have tested on Mac & Linux)
2. Ensure we have ssh access to the ROS system

### Clone this repo

```shell
git clone https://github.com/shawnbmccarthy/viam-ros-module
```

**Note**: if git is not installed, `sudo apt install git`

### Setup Go

Next we will need to install Go on the yahboom robot, we will assume jetson (the instructions will be similar for a Raspberry pi):
1. Log into ROS robot: `ssh <user>@<machine.local>`
2. Download Go: `curl -OL https://golang.org/dl/go1.20.6.linux-aarch64.tar.gz`, please make sure to download the appropriate architecture version
3. Validate the integrity: `sha256sum <go install file>` and compare to checksum on [site](https://go.dev/dl/)
4. Now install: `sudo tar -C /usr/local -xvf go1.20.6.linux-aarch64.tar.gz`
5. Setup Go paths: add this to your profile: `export PATH=$PATH:/usr/local/go/bin`

Now we can test the go install:
```shell
source ~/.profile
go version
```
The output should be something like:
```shell
go version go1.20.6 linux/arm64
```

### Build ROS Module

Once we confirmed go is installed, we must build or rosmodule binary.

```shell
cd <path-to>/viam-ros-module
go build -o rosmodule ./cmd/module/cmd.go
```

This will build our rosmodule binary which will be used by the viam config to start the [modular resources](https://docs.viam.com/extend/modular-resources/)
If we use a different binary name for the rosmodule this name will need to be updated in the config step below.

### Install Viam

To create a viam robot, to set up viam follow the [install instructions](https://docs.viam.com/installation/).

**Note**: Viam requires `libfuse2` to run the appImage, be sure to install the library: `sudo apt install libfuse2`


## Configure Viam
Once viam is installed and running (`sudo systemctl status viam-server.service`), we can use the config found in 
[viam_config.json](./example_viam_config/viam_config.json) to configure viam to work with ros.

**Note**: The configuration contains a network change for the local viam webserver to run on port `9090` as our ROS robot was
already configured to use the `8080` port for its local ROS bridge webserver.

This config contains:
1. components for: cmd_vel topic, image topic, lidar topic, imu topic, battery topic, and edition topic.
2. Attributes for ROS: Validate the topics as well as that roscore is running and the port used (we left the default port of 11311).
3. Modular component: Validate binary name

Once the config is validated, select save and the config will be deployed to the robot, you can now go to the control tab to
test out the robot.

# Contributions

All Pull requests and issues are welcome. The team will track and implement as required. If there are any issues you can email:

* [shawn@viam.com](mailto:shawn@viam.com)
* [solution-eng@viam.com](mailto:solution-eng@viam.com)

# Topics found on Yahboom ROS Robot

Below are the list of topics found on the Yahboom robot, over time we will implement more of the robot features for 
demonstration purposes

## Topics Implemented

* /edition [transbot_msgs/Edition] 1 publisher
* /cmd_vel [geometry_msgs/Twist] 3 publishers
* /image [sensor_msgs/Image] 1 publisher
* /voltage [transbot_msgs/Battery] 1 publisher
* /transbot/imu [sensor_msgs/Imu] 1 publisher
* /scan [sensor_msgs/LaserScan] 1 subscriber

## Topics to Implement

* /control_mode [std_msgs/Int32] 1 subscriber
* /tf [tf2_msgs/TFMessage] 2 subscribers
* /PWMServo [transbot_msgs/PWMServo] 1 subscriber
* /transbot/get_vel [geometry_msgs/Twist] 1 subscriber
* /set_pose [geometry_msgs/PoseWithCovarianceStamped] 1 subscriber
* /TransbotPatrol/parameter_descriptions [dynamic_reconfigure/ConfigDescription] 1 subscriber
* /tf_static [tf2_msgs/TFMessage] 2 subscribers
* /PatrolWarning [transbot_msgs/PatrolWarning] 1 subscriber
* /Adjust [transbot_msgs/Adjust] 1 subscriber
* /TransbotPatrol/parameter_updates [dynamic_reconfigure/Config] 1 subscriber
* /rosout [rosgraph_msgs/Log] 1 subscriber
* /joint_states [sensor_msgs/JointState] 1 subscriber
* /odom_raw [nav_msgs/Odometry] 1 subscriber
* /TargetAngle [transbot_msgs/Arm] 1 subscriber
* /image_msg [transbot_msgs/Image_Msg] 2 publishers
* /PWMServo [transbot_msgs/PWMServo] 3 publishers
* /odom [nav_msgs/Odometry] 1 publisher
* /edge_detection/image/compressedDepth [sensor_msgs/CompressedImage] 1 publisher
* /Graphics_topic [transbot_msgs/General] 2 publishers
* /colorTracker [transbot_msgs/General] 2 publishers
* /transbot/get_vel [geometry_msgs/Twist] 1 publisher
* /contour_moments/moments [opencv_apps/MomentArrayStamped] 1 publisher
* /edge_detection/image/compressed [sensor_msgs/CompressedImage] 1 publisher
* /Autopilot [transbot_msgs/General] 2 publishers
* /diagnostics [diagnostic_msgs/DiagnosticArray] 2 publishers
* /edge_detection/image [sensor_msgs/Image] 1 publisher
* /rosout_agg [rosgraph_msgs/Log] 1 publisher
