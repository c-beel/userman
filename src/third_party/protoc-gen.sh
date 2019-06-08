#!/bin/bash

protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api/v1 api/proto/v1/userman.proto
protoc-go-inject-tag --input=pkg/api/v1/userman.pb.go --XXX_skip gorm
