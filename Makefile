ifneq ("$(wildcard .env)","")
  $(info using .env)
  include .env
endif
# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## test: test all code
.PHONY: test
test:
	go test -race -vet=off -coverpkg ./... -v -coverprofile=cover.out  ./...
	go tool cover -func=cover.out

# test: local test all code
.PHONY: local_test
local_test:
	go test -race -vet=off -coverpkg  ./... -v -coverprofile=cover.out  ./...
	go tool cover -html=cover.out

.PHONY: audit
audit:
	@echo 'Formatting code...'
	@docker run \
		--rm -t \
		-v "$(shell pwd)/tmp/golangci-cache:/.cache" \
		-v "$(shell pwd):/app" \
		--workdir /app \
		golangci/golangci-lint \
		golangci-lint run --fix --config "/app/.golangci.yml" --verbose

## tidy: tidy dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@docker run \
		--rm -t \
		-v "$(shell pwd)/tmp/golangci-cache:/.cache" \
		-v "$(shell pwd):/app" \
		--workdir /app \
		golangci/golangci-lint:latest \
		golangci-lint run --fix --verbose

## sonar: local sonar qube analysis
.PHONY: sonar --sonar-container-create --sonar-network-create sonar-prune
sonar: test --sonar-network-create --sonar-container-create
	@echo 'waiting for sonarqube container healthy status...'
	@until (docker inspect -f {{.State.Health.Status}} sonar | grep -w 'healthy'); do \
		if (docker inspect -f {{.State.Health.Status}} sonar | grep -w 'unhealthy'); then \
			echo 'sonarqube container unhealthy... exiting'; \
			exit 1; \
		fi; \
		sleep 1; \
	done;

	@echo 'sonarqube container healthy! starting scan'
	@docker run \
	     --rm \
         -e SONAR_HOST_URL="${LOCAL_SONAQUBE_ADDR}" \
         -e SONAR_LOGIN="${LOCAL_SONAQUBE_LOGIN}" \
         -e SONAR_PASSWORD="${LOCAL_SONAQUBE_PASSWORD}" \
         -v "$(shell pwd):/usr/src" \
         --network=sonar-net \
         sonarsource/sonar-scanner-cli

--sonar-container-create:
	@if (docker container ls | grep 'sonar'); then \
		echo 'sonarqube container already running.. connecting'; \
	else \
	  	echo 'creating sonarqube container'; \
		docker run \
			-d \
			--rm \
			-p 9000:9000 \
			--expose=9000 \
			--name=sonar \
			--network=sonar-net \
			--health-cmd="wget -qO- http://localhost:9000/api/system/status | grep -w 'UP' && exit 0" \
			--health-interval=4s \
			--health-retries=10 \
 			--health-start-period=30s \
			--health-timeout=5s \
			sonarqube:latest; \
  	fi

--sonar-network-create:
	@if (docker network ls | grep 'sonar-net'); then \
		echo 'sonarqube network already up.. connecting'; \
	else \
	  	echo 'creating sonarqube network'; \
		docker network create sonar-net; \
  	fi

sonar-prune:
	@docker container stop sonar
	@docker network rm sonar-net

# ==================================================================================== #
# DEV
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/api/...

current_time = $(shell date "+%Y-%m-%dT%H:%M:%S%z")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api

# ==================================================================================== #
# DOCKER
# ==================================================================================== #
## up: initiate docker containers
.PHONY: up
up:
	docker compose up --build -d

## down: removes docker containers
.PHONY: down
down:
	docker compose down

## tail: tail consumers logs
.PHONY: api-logs
api-logs:
	docker compose logs api -f