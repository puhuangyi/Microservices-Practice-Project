FROM alpine

ADD middle-layer /tmp
ADD ./config/ /tmp/config
ADD start.sh /tmp

RUN mkdir -p /home/middle-layer/config && mkdir -p /home/middle-layer/log && mkdir -p /home/middle-layer/programm && mv /tmp/middle-layer /home/middle-layer/programm && mv /tmp/start.sh /home/middle-layer && mkdir -p /home/middle-layer/config && mv /tmp/config/* /home/middle-layer/config/ && chmod 777 /home/middle-layer/programm/middle-layer && chmod 777 /home/middle-layer/start.sh

EXPOSE 8080

CMD /home/middle-layer/start.sh

