# TODO

This release is considered a beta release of our ROS integration. There are more items we are looking to implement.

## Items to Implement
1. **ROS Base**: currently our ros base will continuously publish `Twist` messages to the `/cmd_vel` topic. For some robots
this will not be needed. We will need to add configurations to support this.
2. **ROS Services**: We will be implementing various ROS services supported by the Yahboom robot to show references of how
we can implement services and easily map these to viam components
3. **ROS Deployment Examples**: Using our modular registry deployment, we can deploy code required to the robot.
4. **Testing**: Introduce more testing for the components we have built
5. **Building**: Cleanup build documentation, include cross compile and demo dev environment requirements
6. **Documentation**: Clean up documentation and validate build process based on modular registry documentation

## Contributing

All Pull requests and issues are welcome. The team will track and implement as required. If there are any issues you can email:

* [shawn@viam.com](mailto:shawn@viam.com)
* [solution-eng@viam.com](mailto:solution-eng@viam.com)