#!/bin/bash
rm -rf proto
protoc --proto_path=./ --go_out=./ --go-grpc_out=./  ./*.proto
