@echo off

DEL /S "proto"
protoc --proto_path=./ "--go_out=./" "--go-grpc_out=./" "./*.proto"