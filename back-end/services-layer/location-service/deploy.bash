#!/bin/bash

echo "Start build location service"

bash_path=$(dirname $0)

cd $bash_path

go build --tags netgo -o ./location-service ./main/main.go

echo "Build go code successful"

sudo docker build -f Dockerfile -t location-service:1.0 .

cd -

echo "Build location service successful"