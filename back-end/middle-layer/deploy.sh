#!/bin/bash

echo "Start build middle layer"

bash_path=$(dirname $0)

cd $bash_path

go build --tags netgo -o ./middle-layer ./main/main.go

echo "Build go code successful"

sudo docker build -f Dockerfile -t middle-layer:1.0 .

cd -

echo "Build middle layer successful"