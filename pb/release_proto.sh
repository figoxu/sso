#!/usr/bin/env bash
protoc -I . --go_out=plugins=grpc:./sso ./api.proto
protoc -I . --java_out=./java ./api.proto
protoc --doc_out=./doc --doc_opt=html,index.html ./api.proto
protoc --doc_out=./doc --doc_opt=markdown,README.md ./api.proto