# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

# docker build -t wof-roundhoused .
# docker run -p 6161:8080 -e BASE='https://example.com/' -e HOST='0.0.0.0' wof-roundhoused
# curl localhost:6161/1234
# <a href="https://example.com/123/4/1234.geojson">See Other</a>.

FROM golang

ADD . /go-whosonfirst-roundhouse

RUN cd /go-whosonfirst-roundhouse; make bin

EXPOSE 8080

CMD /go-whosonfirst-roundhouse/bin/wof-roundhoused -host ${HOST} -base ${BASE}
