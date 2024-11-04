# Define variables
DOCKER_COMPOSE = docker compose

# Targets
.PHONY: up down restart logs db-shell

# Start services and build docker images with Docker Compose
up:
	sudo $(DOCKER_COMPOSE) up -d

# Stop services, remove containers and images
down:
	sudo $(DOCKER_COMPOSE) down
	sudo docker image prune -f
	sudo docker rmi crud_app

# Restart services
restart: down up

# Access the PostgreSQL shell
db-shell:
	sudo $(DOCKER_COMPOSE) exec postgres_db psql -U postgres

# View logs
logs:
	sudo $(DOCKER_COMPOSE) logs -f
