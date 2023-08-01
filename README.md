# viam-yahboom-transbot-ros

An example of using a [Yahboom ROS Robot](https://category.yahboom.net/collections/ros-robotics/products/transbot-jetson_nano) with [Viam](https://www.viam.com/) integrated.

## About
The ROS robot is built with [ROS Melodic](http://wiki.ros.org/melodic), since this an older release of ROS which uses `Python3.6`
we make use of [goroslib](https://github.com/bluenviron/goroslib) which is a go implementation of the ROS library. It supports:
1. Node Creation
2. Publish/Subscribe
3. Services
4. Actions
5. Standard & Custom Messages

## Requirements
ROS robot with Melodic or Noetic, the package here only supports ROS, ROS2 will be released later.
A [Viam account](https://app.viam.com) with an Organization & Location setup

We currently support the following ROS types:
1. [Twist Msgs](http://docs.ros.org/en/noetic/api/geometry_msgs/html/msg/Twist.html): These will map to the Viam [Base](https://github.com/viamrobotics/rdk/blob/main/components/base/base.go) component
2. [Image](http://docs.ros.org/en/noetic/api/sensor_msgs/html/msg/Image.html): The current codebase supports GRB8 encoding which maps to Viam [Camera Component & VideoSource interface](https://github.com/viamrobotics/rdk/blob/main/components/camera/camera.go)
3. [LaserScan](http://docs.ros.org/en/melodic/api/sensor_msgs/html/msg/LaserScan.html): These will map to the Viam [PointCloud](https://github.com/viamrobotics/rdk/blob/main/pointcloud/pointcloud.go) type by making use of the [Camera Component](https://github.com/viamrobotics/rdk/blob/main/components/camera/camera.go)
4. [Imu](http://docs.ros.org/en/noetic/api/sensor_msgs/html/msg/Imu.html): These will map to the Viam [MovementSensor](https://github.com/viamrobotics/rdk/blob/main/components/movementsensor/movementsensor.go) component
5. [Custom Sensor Msgs](http://www.yahboom.net/study/Transbot-jetson_nano): These custom messages are taken from the Yahboom workspace and coverted to Viam [Sensor](https://github.com/viamrobotics/rdk/blob/main/components/sensor/sensor.go) components

## Installation
We will assume your Robot with ROS is already running and all nodes to support the robot are up and running

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
5. Setup Go paths: add this to your profile - `export PATH=$PATH:/usr/local/go/bin`

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
Once viam is installed and running (`sudo systemctl status viam-server.service`), we can use the config found in `<path-to>/viam-ros-module/viam_config.json`
to configure viam to work with ros.

**Note**: The configuration contains a network change for the local viam webserver to run on port `9090` as our ROS robot was
already configured to 

This config contains:
1. components for: cmd_vel topic, image topic, lidar topic, imu topic, battery topic, and edition topic. 
2. Attributes for ROS: Validate the topics as well as that roscore is running and the port used (we left the default port of 11311).
3. Modular component: Validate binary name

Once config is validate, select save and the config will be deployed to the robot, you can now go to the control tab to 
test out the robot.

# Contribute or Need help

If you would like to contribute, pull requests are welcome

If you need help, you can raise issues

# TODO

1. testing 
2. validate reconfig logic
3. rename ros base (trackedbase should be more generic)
4. Understand and integrate more of the ROS robot as needed

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