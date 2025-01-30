include .env # read from .env file

OUTPUT_BINARY=mongo-crud
OUTPUT_DIR=./bin
ENTRY_DIR = ./cmd/api

.PHONY: build
build:
	@echo "Building..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY) $(ENTRY_DIR)

.PHONY: run
run: build
	@$(OUTPUT_DIR)/$(OUTPUT_BINARY)

.PHONY: clean
clean:
	@rm -rf $(OUTPUT_DIR)

# docker commands
.PHONY: up
up:	
	@echo "Starting containers..."
	@docker compose up -d

.PHONY: down
down:
	@echo "Stopping containers and deleting volumes..."
	@docker compose down -v

.PHONY: list-con
list-containers:
	docker container ls


# Seeding with users
SCRIPTS_DIR=./scripts

.PHONY: seed-db
seed-db:
	@echo "Seeding database with users..."
	bash $(SCRIPTS_DIR)/seed_users.sh


.PHONY: restart
restart: down up run
	@echo "Restarting..."
