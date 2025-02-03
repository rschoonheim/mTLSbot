#!/bin/bash

# Build the root_server image
#
docker build --network=host -t mtls-bot-root-server:latest -f .docker/root_server/Dockerfile .