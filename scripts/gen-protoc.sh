#!/bin/bash

# Generate the protobuffers
protoc --go_out=. \
    --go-grpc_out=. \
    api/*.proto
