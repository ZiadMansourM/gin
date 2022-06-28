THIS_FILE := $(lastword $(MAKEFILE_LIST))
dev_compose_file = docker-compose-dev.yml
staging_compose_file = docker-compose-staging.yml
production_compose_file = docker-compose-production.yml
.PHONY: help build up start down destroy stop restart logs logs-api ps login-timescale login-api db-shell
help:
	make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
production:
	docker-compose -f $(production_compose_file) up $(c)
staging:
	docker-compose -f $(staging_compose_file) build $(c)
build:
	docker-compose -f $(dev_compose_file) build $(c)
up:
	docker-compose -f $(dev_compose_file) up -d $(c)
reload:
	docker-compose -f $(dev_compose_file) down $(c)
	docker-compose -f $(dev_compose_file) up -d $(c)
ps:
	docker-compose -f $(dev_compose_file) ps
update:
	docker-compose -f $(dev_compose_file) up --build -d $(c)
start:
	docker-compose -f $(dev_compose_file) start $(c)
down:
	docker-compose -f $(dev_compose_file) down $(c)
destroy:
	docker-compose -f $(dev_compose_file) down -v $(c)
stop:
	docker-compose -f $(dev_compose_file) stop $(c)
restart:
	docker-compose -f $(dev_compose_file) stop $(c)
	docker-compose -f $(dev_compose_file) up -d $(c)
logs:
	docker-compose -f $(dev_compose_file) logs --tail=100 -f $(c)
logs-api:
	docker-compose -f $(dev_compose_file) logs --tail=100 -f api
login-timescale:
	docker-compose -f $(dev_compose_file) exec timescale /bin/bash
login-api:
	docker-compose -f $(dev_compose_file) exec api /bin/bash
db-shell:
	docker-compose -f $(dev_compose_file) exec timescale psql -Upostgres