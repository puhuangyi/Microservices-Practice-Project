#!/bin/bash

echo "Start build payment service"

bash_path=$(dirname $0)

cd $bash_path

sudo go build --tags netgo -o ./payment-service ./main/main.go

echo "Build go code successful"

sudo docker build -f Dockerfile -t payment-service:1.0 .

cd -

echo "Build payment service successful"