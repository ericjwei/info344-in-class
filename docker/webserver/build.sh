#!/usr/bin/env bash
set -e #stops if there is an error
GOOS=linux go build
docker build - t "dockerhub-name"/testserver .
docker push  "dockerhub-name"/testserver
