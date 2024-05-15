#!/bin/bash

trap "exit 1" ERR
cp ./assets/icons/advertising128.png ./internal/front/ui/gtk/application/objectmaker/advertisementer128.png
GOARCH=amd64 go build -C ./cmd/advertisementer/ -v -ldflags "-w -s -linkmode=external -X main.version=$2" -o $1
echo "====> building for linux is done <===="