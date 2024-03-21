#!/bin/bash
echo "go build"
go mod tidy
go build -o go-kuaiyin main.go
chmod +x ./go-kuaiyin
echo "kill go-kuaiyin service"
killall go-kuaiyin # kill go-admin service
nohup ./go-kuaiyin server -c=config/settings.dev.yml -a=true >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
echo "run go-kuaiyin success"
ps -aux | grep go-kuaiyin
tail -f access.log