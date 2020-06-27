FROM golang:1.13-alpine
ARG BACKEND_HOST=localhost

WORKDIR /go/src/app
COPY . .

RUN go install ./...