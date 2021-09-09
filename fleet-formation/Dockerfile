FROM golang:1.13-alpine AS base
RUN apk add build-base

FROM base AS build-base
WORKDIR /go/src/app
COPY fleet-formation ./fleet-formation
COPY skysign-common ./skysign-common
COPY skysign-proto ./skysign-proto

FROM build-base AS dependency
WORKDIR /go/src/app/fleet-formation
RUN make dependency

FROM dependency AS test
RUN make unit-test

FROM dependency AS build
RUN make install

FROM alpine:latest
WORKDIR /app/

EXPOSE 5000

COPY --from=build /go/bin/fleet-formation .