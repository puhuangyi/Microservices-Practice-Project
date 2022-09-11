#!/bin/sh

cd /home/location-service

nohup /home/location-service/location-service >> /home/location-service/log/nohup.log 2>&1 &

tail -f /home/location-service/log/nohup.log