FROM golang:1.13-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go install ./...

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /go/bin/skysign-edge .