#!/bin/bash

echo "Start build login service"

bash_path=$(dirname $0)

cd $bash_path

sudo go build --tags netgo -o ./login-service ./main/main.go

echo "Build go code successful"

sudo docker build -f Dockerfile -t login-service:1.0 .

cd -

echo "Build login service successful"