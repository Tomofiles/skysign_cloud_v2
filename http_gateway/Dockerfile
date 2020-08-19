FROM golang:1.13-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go install ./...

FROM alpine:latest

WORKDIR /app/

EXPOSE 5000

COPY --from=builder /go/bin/vehicle-gateway .
COPY --from=builder /go/bin/mission-gateway .
COPY --from=builder /go/bin/helper-gateway .
COPY --from=builder /go/bin/comm-gateway .