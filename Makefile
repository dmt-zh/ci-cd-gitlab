ifeq (,$(wildcard .env))
$(error .env file is missing. Please create one based on .env.example. Run: "cp .env.example .env" and fill in the missing values.)
endif

include .env

# --- Running commands ---
init: # Create credentials in the env file
	@bash setup.sh

start: # Start all services
	@docker compose up -d

stop: # Stop all services
	@docker compose down --remove-orphans

ps: # List of running containers and their status
	@docker compose ps --format "table {{.Name}}\t{{.Status}}"
