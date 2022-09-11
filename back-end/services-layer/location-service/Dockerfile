FROM alpine

ADD location-service /tmp
ADD ./config/ /tmp/config
ADD start.sh /tmp

RUN mkdir -p /home/location-service/config && mkdir -p /home/location-service/log && mv /tmp/location-service /home/location-service/ && mv /tmp/start.sh /home/location-service && mv /tmp/config/* /home/location-service/config && chmod 777 /home/location-service/location-service && chmod 777 /home/location-service/start.sh

EXPOSE 9000

CMD /home/location-service/start.sh