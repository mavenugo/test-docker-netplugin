FROM golang:1.6.3-alpine

RUN apk update

RUN mkdir -p /run/docker/plugins

COPY . /go/src/github.com/mavenugo/test-docker-netplugin
WORKDIR /go/src/github.com/mavenugo/test-docker-netplugin

RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
    gcc libc-dev \
    && go install --ldflags '-extldflags "-static"' \
    && apk del .build-deps

CMD ["/go/bin/test-docker-netplugin"]
