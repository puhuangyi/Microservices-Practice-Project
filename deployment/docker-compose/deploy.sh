#!/bin/bash

bash_path=$(dirname $0)

cd $bash_path

echo "now path is $bash_path"

docker build -f Dockerfile -t query-service:1.0 .

chmod 775 $bash_path/../../back-end/services-layer/payment-service/deploy.sh
$bash_path/../../back-end/services-layer/payment-service/deploy.sh

chmod 775 $bash_path/../../back-end/services-layer/location-service/deploy.sh
$bash_path/../../back-end/services-layer/location-service/deploy.sh

chmod 775 $bash_path/../../back-end/services-layer/login-service/deploy.sh
$bash_path/../../back-end/services-layer/login-service/deploy.sh

chmod 775 $bash_path/../../back-end/middle-layer/deploy.sh
$bash_path/../../back-end/middle-layer/deploy.sh