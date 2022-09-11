#!/bin/sh

cd /home/login-service

nohup /home/login-service/login-service >> /home/login-service/log/nohup.log 2>&1 &

tail -f /home/login-service/log/nohup.log