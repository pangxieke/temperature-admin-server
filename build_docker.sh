#!/bin/bash

set -e

if [ $# -lt 2 ];then
    echo "use: build_docker.sh versionNum dev[release] "
    echo "example: build_docker.sh dev  v0.0.1 "
    exit 1
fi

#git_commitid=$(git describe --always --tags)
#commitid="_$1.$2_${git_commitid}_"`date +%Y%m%d%H%M%S`
commitid="_$1.$2_"`date +%Y%m%d%H%M%S`
#echo "commitid = ${commitid}"

host="registry.cn-shenzhen.aliyuncs.com/***"

app="temp_back"
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app_d ./main.go

docker build -t ${host}/${app}:${commitid} .

#docker push ${host}/${app}:${commitid}

echo "${host}/${app}:${commitid}"
