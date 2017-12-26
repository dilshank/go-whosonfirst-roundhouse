# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

# docker build -t wof-roundhoused .
# docker run -it -p 6161:8080 -e BASE='https://example.com/' -e HOST='0.0.0.0' wof-roundhoused
# curl localhost:6161/1234
# <a href="https://example.com/123/4/1234.geojson">See Other</a>.

# build phase - see also:
# https://medium.com/travis-on-docker/multi-stage-docker-builds-for-creating-tiny-go-images-e0e1867efe5a
# https://medium.com/travis-on-docker/triple-stage-docker-builds-with-go-and-angular-1b7d2006cb88

FROM golang:alpine AS build-env

# https://github.com/gliderlabs/docker-alpine/issues/24

RUN apk add --update alpine-sdk

ADD . /go-whosonfirst-roundhouse

RUN cd /go-whosonfirst-roundhouse; make bin

FROM alpine

WORKDIR /go-whosonfirst-static/bin/

COPY --from=build-env /go-whosonfirst-roundhouse/bin/wof-roundhoused /go-whosonfirst-roundhouse/bin/wof-roundhoused

EXPOSE 8080

CMD /go-whosonfirst-roundhouse/bin/wof-roundhoused -host ${HOST} -base ${BASE}
