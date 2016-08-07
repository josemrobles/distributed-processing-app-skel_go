FROM golang:1.6

WORKDIR /go/src/github.com/josemrobles/distributed-processing-app-skel_go/api

RUN apt-get update &&\
    apt-get install -y supervisor ;\
    mkdir -p /var/log/supervisor 

COPY supervisor.conf /etc/supervisor/conf.d/supervisor.conf
ADD . /go/src/github.com/josemrobles/distributed-processing-app-skel_go/api
CMD "./entrypoint.sh"
