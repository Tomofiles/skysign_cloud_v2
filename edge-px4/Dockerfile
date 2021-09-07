FROM golang:1.13-alpine AS base
RUN apk add build-base

FROM base AS build-base
WORKDIR /go/src/app
COPY . .

FROM build-base AS dependency
RUN make dependency

# FROM dependency AS test
# RUN make unit-test

FROM dependency AS build
RUN make install

FROM alpine:latest
WORKDIR /app/

COPY --from=build /go/bin/edge-px4 .