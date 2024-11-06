# Define variables
DOCKER_COMPOSE = sudo docker compose
DOCKER_COMPOSE_W = docker-compose
DOCKER = sudo docker
DOCKER_W = docker
POSTGRES_VERSION = 13

# Targets
.PHONY: up down restart logs db-shell

# Start services and build docker images with Docker Compose
up:
	$(DOCKER_COMPOSE_W) up -d


# Stop services, remove containers and images
down:
	$(DOCKER_COMPOSE_W) down
	$(DOCKER_W) image prune -f
	$(DOCKER_W) rmi crud_app
	$(DOCKER_W) rmi postgres:$(POSTGRES_VERSION)

# Restart services
restart: down up

# Access the PostgreSQL shell
db-shell:
	$(DOCKER_COMPOSE_W) exec postgres_db psql -U postgres

# View logs
logs:
	$(DOCKER_COMPOSE_W) logs -f

#for linux users
upl:
	$(DOCKER_COMPOSE) up -d

downl:
	$(DOCKER_COMPOSE) down
	$(DOCKER) image prune -f
	$(DOCKER) rmi crud_app
	$(DOCKER) rmi postgres:$(POSTGRES_VERSION)

restartl: downl upl

db-shelll:
	$(DOCKER_COMPOSE) exec postgres_db psql -U postgres

logsl:
	$(DOCKER_COMPOSE) logs -f