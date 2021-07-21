# https://hub.docker.com/r/hypnoglow/golang-migrate/dockerfile
# https://github.com/docker-library/golang/blob/906e04de73168f643c5c2b40dca0877a14d2377c/1.10/alpine3.7/Dockerfile
FROM golang:1.15-alpine AS build

ENV MIGRATE_VERSION v3.2.0

RUN apk add --no-cache git

# Migrate cli
RUN go get -v -d github.com/golang-migrate/migrate/cli \
    && go get -v -d github.com/lib/pq

WORKDIR /go/src/github.com/golang-migrate/migrate

RUN git checkout ${MIGRATE_VERSION} \
    && go build -tags 'postgres' -o ./bin/migrate ./cli

# Swagger cli
RUN go get -u -d github.com/swaggo/swag/cmd/swag

WORKDIR /go/src/github.com/swaggo/swag/cmd

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./bin/swag ./swag

FROM golang:1.16-alpine

ARG USER_ID
ARG GROUP_ID

RUN apk add --no-cache gcc g++ make bash coreutils unzip

RUN if [ ${USER_ID:-0} -ne 0 ] && [ ${GROUP_ID:-0} -ne 0 ]; then \
    if getent passwd www-data ; then deluser www-data; fi &&\
    if getent group www-data ; then delgroup www-data; fi &&\
    addgroup -g ${GROUP_ID} www-data &&\
    adduser -u ${USER_ID} -D -S -G www-data www-data &&\
    install -d -m 0755 -o www-data -g www-data /home/www-data\
;fi

COPY ./common/wait-for-it.sh /usr/local/bin/wait-for-it
RUN chmod 555 /usr/local/bin/wait-for-it

COPY --from=build /go/src/github.com/golang-migrate/migrate/bin/migrate /usr/local/bin/migrate
COPY --from=build /go/src/github.com/swaggo/swag/cmd/bin/swag /usr/local/bin/swag

WORKDIR /app

USER www-data
