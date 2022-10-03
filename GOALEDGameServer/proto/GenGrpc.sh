#!/bin/sh

protoc --go_out=../go/pb --go_opt=paths=source_relative --go-grpc_out=../go/pb --go-grpc_opt=paths=source_relative Multiplay.proto

protoc --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc="/grpc_plugins/grpc_csharp_plugin" Multiplay.proto
