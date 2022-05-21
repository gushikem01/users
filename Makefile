HAS_YQ=$(shell which yq)
HAS_MIGRATE=$(shell which migrate)
NS=session

.PHONY: up
up:	 ## docker-compose up -d
	docker-compose up -d

.PHONY: build
build: ## docker-compose build
	docker-compose build

.PHONY: down
down: ## docker-compose down
	docker-compose down

# air api
.PHONY: air-%
air-%: ## air -c $(@:air-%=%)/cmd/api/.air.toml
	cd ./backend/go && air -c $(@:air-%=%)/cmd/api/.air.toml

# task
.PHONY: task
task:
	cd task/cmd/cli && go run main.go
