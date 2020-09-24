# This Dockerfile is used to run the server

FROM golang:1.14-alpine AS builder

WORKDIR go/src/arise

COPY . .

WORKDIR server

RUN go build \
    && mv server /go/bin


#####
FROM alpine:3.6

COPY --from=builder /go/bin/server /usr/local/bin

ENTRYPOINT ["server"]
