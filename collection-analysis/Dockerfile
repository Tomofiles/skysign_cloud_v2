FROM golang:1.13-alpine AS base
RUN apk add build-base

FROM base AS dependency-base
WORKDIR /go/src/app
COPY collection-analysis/go.* ./collection-analysis/
COPY collection-analysis/Makefile ./collection-analysis/
COPY skysign-common/go.* ./skysign-common/
COPY skysign-common/Makefile ./skysign-common/
COPY skysign-proto/go.* ./skysign-proto/
COPY skysign-proto/Makefile ./skysign-proto/

FROM dependency-base AS dependency
WORKDIR /go/src/app/collection-analysis
RUN make dependency

FROM dependency AS build-base
WORKDIR /go/src/app
COPY collection-analysis ./collection-analysis
COPY skysign-common ./skysign-common
COPY skysign-proto ./skysign-proto

FROM build-base AS test
WORKDIR /go/src/app/collection-analysis
RUN make unit-test

FROM build-base AS build
WORKDIR /go/src/app/collection-analysis
RUN make install

FROM alpine:latest
WORKDIR /app/

EXPOSE 5000

COPY --from=build /go/bin/collection-analysis .