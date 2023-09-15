#!/usr/bin/env bash

#
# rosmodule.sh
# wrapper to the go executable
# This file can be customized to:
# (1) deliver code
# (2) build required execution environment
# (3) with go - create a simple internal web server to get various ROS statuses
#
# When creating scripts to create code it is important to introduce techniques to
# ensure that we execute that section of the script only once per version upgrade
#

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
#VERSION_FILE="./.version"

#
# custom environment actions
#
source /opt/ros/melodic/local_setup.bash --extend
source /home/jetson/transbot_ws/setup.bash --extend

#
# deploy docker
#
exec ${SCRIPT_DIR}/rosmodule