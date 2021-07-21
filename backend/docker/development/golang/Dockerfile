FROM golang:1.16-alpine

RUN apk add --no-cache curl gcc g++ make

RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /app

COPY ./development/golang/entrypoint.sh /usr/local/bin/docker-golang-entrypoint
RUN chmod +x /usr/local/bin/docker-golang-entrypoint

HEALTHCHECK --interval=5s --timeout=3s --start-period=1s \
    CMD curl --fail http://127.0.0.1:9000/health || exit 1
