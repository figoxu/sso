#!/usr/bin/env bash
protoc --go_out=./sso --java_out=./java ./api.proto
protoc --doc_out=./doc --doc_opt=html,index.html ./api.proto
protoc --doc_out=./doc --doc_opt=markdown,README.md ./api.proto