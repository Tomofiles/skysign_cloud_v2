FROM golang:1.13-alpine AS base
RUN apk add build-base

FROM base AS dependency-base
WORKDIR /go/src/app
COPY edge-px4/go.* ./edge-px4/
COPY edge-px4/Makefile ./edge-px4/
COPY skysign-proto/go.* ./skysign-proto/
COPY skysign-proto/Makefile ./skysign-proto/

FROM dependency-base AS dependency
WORKDIR /go/src/app/edge-px4
RUN make dependency

FROM dependency AS build-base
WORKDIR /go/src/app
COPY edge-px4 ./edge-px4
COPY skysign-proto ./skysign-proto

FROM build-base AS test
WORKDIR /go/src/app/edge-px4
RUN make unit-test

FROM build-base AS build
WORKDIR /go/src/app/edge-px4
RUN make install

FROM alpine:latest
WORKDIR /app/

COPY --from=build /go/bin/edge-px4 .