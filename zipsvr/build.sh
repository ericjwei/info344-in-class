#!/usr/bin/env bash
set -e
echo "building go server for Linux..."
GOOS=linux go build 
docker build -t ericjwei/zipsvr .
docker push ericjwei/zipsvr
go clean
#!docker run -d -p 80:80 ericjwei/zipsvr