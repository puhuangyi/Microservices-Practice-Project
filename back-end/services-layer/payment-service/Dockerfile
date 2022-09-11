FROM alpine

ADD payment-service /tmp
ADD ./config/ /tmp/config
ADD start.sh /tmp

RUN mkdir -p /home/payment-service/config && mkdir -p /home/payment-service/log && mv /tmp/payment-service /home/payment-service/ && mv /tmp/start.sh /home/payment-service && mv /tmp/config/* /home/payment-service/config && chmod 777 /home/payment-service/payment-service && chmod 777 /home/payment-service/start.sh

EXPOSE 9000

CMD /home/payment-service/start.sh