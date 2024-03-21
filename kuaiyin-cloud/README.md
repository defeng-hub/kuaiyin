
# 一、安装 
1. protoc 
2. protoc-gen-go 
3. protoc-gen-go-grpc
4. kratos

```cmd
https://github.com/protocolbuffers/protobuf/releases

https://github.com/protocolbuffers/protobuf/releases/tag/v21.6
```

```cmd
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

$ protoc-gen-go --version
protoc-gen-go v1.32.0
$ protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.3.0
```


```cmd
go install  github.com/go-kratos/kratos/cmd/kratos/v2@latest

$ kratos -v
```

# 二、尝试一个服务

```cmd
$ kratos new verify-code
$ cd verify-code
$ go mod tidy
$ go get github.com/google/wire/cmd/wire
$ go generate ./...
$ kratos run
```

# 常用命令
```cmd
创建proto
$ kratos proto add api/verifyCode/verifyCode.proto

生成代码（用于生成client stub 相关代码）
$ kratos proto client api/verifyCode/verifyCode.proto

生成代码（用于生成server 相关代码，具体逻辑实现）
$ kratos proto server api/verifyCode/verifyCode.proto -t internal/service

```

# 打包
```
$ go build -o ./bin/ ./... 
$ ./bin/verify-code -conf ./configs
```