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

# mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie \
# https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
# && chmod +x ~/.aws-lambda-rie/aws-lambda-rie
# docker run -d -v ~/.aws-lambda-rie:/aws-lambda --entrypoint /aws-lambda/aws-lambda-rie  -p 9000:8080 myfunction:latest /main
