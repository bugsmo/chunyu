GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
# 这是一个条件判断。它检查GOPATH变量是否为空。如果GOPATH为空（即没有被设置），则使用$(error ...)函数来打印一条错误信息并终止make的执行。这条错误信息提示用户需要在运行make命令之前设置环境变量GOPATH。
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

# 这行代码定义了一个名为FAIL_ON_STDOUT的变量，其值是一个awk命令。awk是一个强大的文本处理工具。这个awk命令的作用是打印所有输入行，并在处理完所有输入后检查是否有任何行被打印（即NR > 0，NR是awk的内置变量，表示当前处理的记录数，也就是行数）。如果有任何行被打印，则awk会退出并返回状态码1，这通常表示一个错误。这个FAIL_ON_STDOUT变量可能是为了在后续的Makefile规则中用来检查某个命令的输出，如果命令产生了任何输出（可能是错误或警告），则使构建过程失败。
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"

ifeq ($(OS),Windows_NT)
    IS_WINDOWS:=1
endif

# 这个变量通过git describe --tags --always命令来获取Git仓库中当前HEAD所指向的最近的标签名。如果HEAD不在任何标签上，它会返回一个包含当前提交哈希值的描述。这个命令常用于获取项目的版本信息。
APP_VERSION=$(shell git describe --tags --always)
# 这个变量计算当前目录相对于其上一级目录的路径。它首先获取当前工作目录（`PWD`）的基本名称（即最后的目录名），然后切换到上一级目录，并获取该目录的基本名称。最后，它将这两个基本名称用斜杠（`/`）连接起来，形成相对路径。
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
# 这个变量使用`sed`命令将`APP_RELATIVE_PATH`中的斜杠（`/`）替换为短横线（`-`）。`sed`的`-E`选项启用了扩展的正则表达式语法，`-n`选项表示只打印经过`sed`处理的行。这个命令的效果是将路径转换为适合用作Docker镜像名称或应用程序名称的格式。
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
# 这个变量用于生成Docker镜像的完整名称。它使用`awk`命令来格式化`APP_NAME`。`-F '@'`设置了输入字段分隔符为`@`（虽然在这个上下文中并没有实际用到，因为`APP_NAME`中不包含`@`字符）。然后，`awk`命令打印出由字符串`"kratos-monolithic-demo/"`、`APP_NAME`的值和固定的标签`:0.1.0`组成的字符串。这通常用于在Docker构建和部署过程中指定镜像名称和版本。
APP_DOCKER_IMAGE=$(shell echo $(APP_NAME) |awk -F '@' '{print "bugsmo/" $$0 ":0.1.0"}')


.PHONY: dep tidy vendor build clean docker gen ent wire api openapi run test cover vet lint app

# download dependencies of module
dep:
	go mod download

# update go.mod
tidy:
	go mod tidy

# create vendor
vendor:
	go mod vendor

# build golang application
build:
ifeq ("$(wildcard ./bin/)","")
	mkdir bin
endif
	@go build -ldflags "-X main.Service.Version=$(APP_VERSION)" -o ./bin/ ./...

# clean build files
clean:
	@go clean
	$(if $(IS_WINDOWS), del "coverage.out", rm -f "coverage.out")

# build docker image
docker:
	@docker build -t $(APP_DOCKER_IMAGE) . \
				  -f ../../../.docker/Dockerfile \
				  --build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) GRPC_PORT=9000 REST_PORT=8000

# generate ent code
# 这是一个条件判断，检查./internal/data/ent这个目录是否存在。
# wildcard是一个Makefile函数，用于查找匹配给定模式的所有文件名。
# 如果目录存在，wildcard会返回非空字符串，否则返回空字符串。
# ifneq 是一个条件判断，检查两个参数是否不相等。如果目录存在（即wildcard返回非空字符串），则条件为真。
ent:
ifneq ("$(wildcard ./internal/data/ent)","")
	@go run -mod=mod entgo.io/ent/cmd/ent generate \
				--feature privacy \
				--feature sql/modifier \
				--feature entql \
				--feature sql/upsert \
				./internal/data/ent/schema
endif

# generate wire code
wire:
	go run -mod=mod github.com/google/wire/cmd/wire ./cmd/server

# generate protobuf api go code
api:
	cd ../../../ && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	cd ../../../ && \
	buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml

# run application
run: api openapi
	@go run ./cmd/server -conf ./configs

# run tests
test:
	@go test ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# run static analysis
vet:
	@go vet

# run lint
lint:
	@golangci-lint run

# build service app
app: api wire ent build

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help