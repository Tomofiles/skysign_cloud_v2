FROM golang:1.13-alpine AS base
RUN apk add build-base

FROM base AS dependency-base
WORKDIR /go/src/app
COPY flight-operation/go.* ./flight-operation/
COPY flight-operation/Makefile ./flight-operation/
COPY skysign-common/go.* ./skysign-common/
COPY skysign-common/Makefile ./skysign-common/
COPY skysign-proto/go.* ./skysign-proto/
COPY skysign-proto/Makefile ./skysign-proto/

FROM dependency-base AS dependency
WORKDIR /go/src/app/flight-operation
RUN make dependency

FROM dependency AS build-base
WORKDIR /go/src/app
COPY flight-operation ./flight-operation
COPY skysign-common ./skysign-common
COPY skysign-proto ./skysign-proto

FROM build-base AS test
WORKDIR /go/src/app/flight-operation
RUN make unit-test

FROM build-base AS build
WORKDIR /go/src/app/flight-operation
RUN make install

FROM alpine:latest
WORKDIR /app/

EXPOSE 5000

COPY --from=build /go/bin/flight-operation .