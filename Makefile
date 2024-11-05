# Define variables
DOCKER_COMPOSE = sudo docker compose
DOCKER = sudo docker
POSTGRES_VERSION = 13

# Targets
.PHONY: up down restart logs db-shell

# Start services and build docker images with Docker Compose
up:
	$(DOCKER_COMPOSE) up -d

# Stop services, remove containers and images
down:
	$(DOCKER_COMPOSE) down
	$(DOCKER) image prune -f
	$(DOCKER) rmi crud_app
	$(DOCKER) rmi postgres:$(POSTGRES_VERSION)

# Restart services
restart: down up

# Access the PostgreSQL shell
db-shell:
	$(DOCKER_COMPOSE) exec postgres_db psql -U postgres

# View logs
logs:
	$(DOCKER_COMPOSE) logs -f
