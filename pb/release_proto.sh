#!/usr/bin/env bash
protoc -I . --go_out=plugins=grpc:./sso ./api.proto
protoc -I . --java_out=./java ./api.proto
protoc --plugin=protoc-gen-grpc-java=/Users/xujianhui/develop/golang/gopath/src/github.com/grpc/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java \
  --grpc-java_out=lite:"./java" --proto_path="$DIR_OF_PROTO_FILE" "./api.proto"
protoc --doc_out=./doc --doc_opt=html,index.html ./api.proto
protoc --doc_out=./doc --doc_opt=markdown,README.md ./api.proto