#!/bin/sh

cd /home/payment-service

nohup /home/payment-service/payment-service >> /home/payment-service/log/nohup.log 2>&1 &

tail -f /home/payment-service/log/nohup.log