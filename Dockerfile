# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

FROM golang

ADD . /go-whosonfirst-roundhouse

RUN cd /go-whosonfirst-roundhouse; make bin

EXPOSE 8080

CMD /go-whosonfirst-roundhouse/bin/wof-roundhoused -host '0.0.0.0'
