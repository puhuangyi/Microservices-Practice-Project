FROM alpine

ADD login-service /tmp
ADD ./config/ /tmp/config
ADD start.sh /tmp

RUN mkdir -p /home/login-service/config && mkdir -p /home/login-service/log && mv /tmp/login-service /home/login-service && mv /tmp/start.sh /home/login-service && mv /tmp/config/* /home/login-service/config/ && chmod 777 /home/login-service/login-service && chmod 777 /home/login-service/start.sh

EXPOSE 9000

CMD /home/login-service/start.sh