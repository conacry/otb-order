FROM golang:1.20.8-alpine3.18 AS builder

ARG INTEGRATION_TEST_DISABLE

ENV GOSUMDB=off
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

RUN apk --no-cache add git gcc musl-dev && \
    rm -rf /var/cache/apk/*

COPY . .
RUN go mod vendor
RUN go test ./test/...

RUN go build -o os-order .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /home/app
COPY --from=builder /app/os-order .
EXPOSE 8080
CMD ["./os-order"]
