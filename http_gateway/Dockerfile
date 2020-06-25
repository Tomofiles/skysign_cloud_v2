FROM golang:1.13-alpine
ARG BACKEND_HOST=localhost

WORKDIR /go/src/app
COPY . .

RUN go install ./...

EXPOSE 5000

ENV BACKEND_HOST=${BACKEND_HOST}

CMD ["sh", "-c", "/go/bin/comm-gateway -backend_host ${BACKEND_HOST}"]