SHELL = /bin/bash

#SCRIPT_DIR         = $(shell pwd)/etc/script
#请选择golang版本
BUILD_IMAGE_SERVER  = golang:1.21
#请选择node版本
BUILD_IMAGE_WEB     = node:18
#项目名称
PROJECT_NAME        = github.com/bugsmo/chunyu/server
#配置文件目录
CONFIG_FILE         = config.yaml
#镜像仓库命名空间
IMAGE_NAME          = bugsmo
#镜像地址
REPOSITORY          = registry.cn-guangzhou.aliyuncs.com/${IMAGE_NAME}

ifeq ($(TAGS_OPT),)
TAGS_OPT            = latest
else
endif

#本地环境打包后端
build-server-local:
	@cd server/ && if [ -f "server" ];then rm -rf server; else echo "OK!"; fi \
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy \
	&& go build -ldflags "-B 0x$(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n') -X main.Version=${TAGS_OPT}" -v

