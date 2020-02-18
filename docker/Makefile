include add_help

COMPOSE := @docker-compose
DOCKER = @docker
DC = $(DOCKER) container
DCE := $(DC) exec
BASH_C := /bin/bash -c
IPA_FORMAT := '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
DEP := $(DC) exec tic-tac-toe-api-container go get ./...
ALL_TESTS := $(DC) exec tic-tac-toe-api-container go test ./... -count=1
C_UP = $(COMPOSE) up
C_DOWN = $(COMPOSE) down

## build, compose up, and daemonize
up:
	$(C_UP) --build -d

## compose down
down:
	$(C_DOWN)

## compose down and remove all volumes
down-v:
	$(C_DOWN) -v

## execute bash inside the container
bash:
	$(DC) exec -it tic-tac-toe-api-container /bin/bash

## attach to the container
attach:
	$(DC) attach tic-tac-toe-api-container

## show IP address of the container
ip:
	$(DC) inspect --format=$(IPA_FORMAT) tic-tac-toe-api-container

## install dependencies
dep:
	$(DEP)

## run all tests from container
all-tests:
	$(ALL_TESTS)

## run all tests from container in verbose mode
all-tests-v:
	$(ALL_TESTS) -v

## run all tests from container in coverage mode
all-tests-c:
	$(ALL_TESTS) -cover

## restart the whole stack
restart:
	$(C_DOWN)
	$(C_UP) --build -d
