# This Dockerfile is used to run the server

FROM golang:1.14-alpine AS builder

WORKDIR go/src/arise

COPY . .

RUN go build \
    && mv arise /go/bin


#####
FROM alpine:3.6

COPY --from=builder /go/bin/arise /usr/local/bin

ENTRYPOINT ["arise", "relay"]
