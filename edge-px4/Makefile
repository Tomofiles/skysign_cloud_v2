.PHONY: dependency build-test unit-test install #integration-test docker-up docker-down clear 

dependency:
	@go mod download

build-test: dependency
	@go build -v ./...

unit-test: dependency
	@go test -v -short ./...

# integration-test: docker-up dependency
# 	@go test -v ./...

install: dependency
	@go install -v ./...

# docker-up:
# 	@docker-compose up -d --build

# docker-down:
# 	@docker-compose down

# clear: docker-down