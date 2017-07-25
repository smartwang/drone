#!/bin/bash
# need yum install -y glibc-static before build

set -e
set -x

go get ./...
cd ${GOPATH}/src/github.com/smartwang/drone-ui
git pull
git checkout develop
cd -
go build -ldflags '-extldflags "-static" -X drone/version.VersionDev=build.'${DRONE_BUILD_NUMBER:-$(date +%s)}  -o release/drone-server  github.com/smartwang/drone/cmd/drone-server
CGO_ENABLED=0 go build -o release/drone-agent github.com/smartwang/drone/cmd/drone-agent

docker build -f Dockerfile -t registry.cn-hangzhou.aliyuncs.com/ly_ops/drone-server ../drone
docker build -f Dockerfile.agent -t registry.cn-hangzhou.aliyuncs.com/ly_ops/drone-agent ../drone

echo "push to your registry"
