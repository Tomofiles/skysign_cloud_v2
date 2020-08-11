FROM golang:1.13-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go install ./...

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /go/bin/vehicle-user-gateway .
COPY --from=builder /go/bin/mission-user-gateway .
COPY --from=builder /go/bin/helper-user-gateway .
COPY --from=builder /go/bin/comm-edge-gateway .