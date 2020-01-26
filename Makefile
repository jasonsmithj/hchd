.PHONY: setup stop start restart up down test migrate seed create-test-db


LOCAL_DOCKER_COMPOSE  = ./docker/docker-compose.local.yml
COMPOSE_OPTS        = -f "$(LOCAL_DOCKER_COMPOSE)" -p hchd

go_build:
	./scripts/build

setup: go_build
	docker-compose $(COMPOSE_OPTS) up -d --build

build: go_build
	docker-compose $(COMPOSE_OPTS) build

stop:
	docker-compose $(COMPOSE_OPTS) stop

start:
	docker-compose $(COMPOSE_OPTS) start

restart:
	docker-compose $(COMPOSE_OPTS) restart

up:
	docker-compose $(COMPOSE_OPTS) up -d

down:
	docker-compose $(COMPOSE_OPTS) down

logs:
	docker-compose $(COMPOSE_OPTS) logs -f

reset: down build up
