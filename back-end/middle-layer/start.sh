#!/bin/sh

cd /home/middle-layer

nohup /home/middle-layer/programm/middle-layer >> /home/middle-layer/log/nohup.log 2>&1 &

tail -f /home/middle-layer/log/nohup.log