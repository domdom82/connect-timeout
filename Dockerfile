FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y iptables

ADD connect-timeout /connect-timeout
ADD run.sh /run.sh

CMD /run.sh

EXPOSE 8080