# Builder
FROM golang:1.16-alpine as builder

RUN apk update && apk upgrade
RUN go version

WORKDIR /app

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata ca-certificates

WORKDIR /app

COPY --from=builder ./.bin/app /usr/local/bin
COPY --from=builder ./configs /usr/local/configs/
COPY --from=builder ./templates /usr/local/templates/

CMD ["./app"]
