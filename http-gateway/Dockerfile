FROM golang:1.13-alpine AS base

FROM base AS build-base
WORKDIR /go/src/app
COPY . .

FROM build-base AS dependency
RUN go mod download

# FROM dependency AS test
# RUN go test ./...

FROM dependency AS build
RUN go install ./...

FROM alpine:latest
WORKDIR /app/

EXPOSE 5000

COPY --from=build /go/bin/vehicle-gateway .
COPY --from=build /go/bin/mission-gateway .
COPY --from=build /go/bin/helper-gateway .
COPY --from=build /go/bin/comm-gateway .
COPY --from=build /go/bin/flightplan-gateway .