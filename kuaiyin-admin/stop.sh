#!/bin/bash
killall go-kuaiyin # kill go-admin service
echo "stop go-kuaiyin success"
ps -aux | grep go-kuaiyin