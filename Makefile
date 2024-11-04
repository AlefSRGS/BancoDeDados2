# Define variables
DOCKER_COMPOSE = docker-compose

# Targets
.PHONY: up down restart logs

# Start services and build docker images with Docker Compose
up: build
	$(DOCKER_COMPOSE) up -d

# Stop services, remove containers and images
down:
	$(DOCKER_COMPOSE) down
	docker image prune -f

# Restart services
restart: down up

# View logs
logs:
	$(DOCKER_COMPOSE) logs -f
