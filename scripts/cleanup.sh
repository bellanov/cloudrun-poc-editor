#!/bin/bash
#
# Clean up Docker resources.

sudo docker container prune
sleep 5
sudo docker image prune
