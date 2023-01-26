up:
	docker compose up -d
up-tenhou-scraper:
	docker compose up -d tenhou_scraper
up-log-db:
	docker compose up -d tenhou_log_db
build:
	docker compose build --no-cache
ps:
	docker compose ps
stop:
	docker compose stop
down:
	docker compose down
down-log-db:
	docker compose down tenhou_log_db
down-all:
	docker-compose down --rmi local --volumes --remove-orphans
exec-terraform:
	docker compose exec terraform bash
log-db:
	docker compose logs tenhou_log_db
setup:
	make build up setup-go setup-terraform
setup-go:
	go mod download
setup-terraform:
	docker compose exec terraform gcloud auth application-default login
	docker compose exec terraform gcloud auth login
	make project-set
	make init-dev
project-set:
	docker compose exec terraform gcloud config set project mj-log-dev
format-terraform:
	docker compose exec terraform terraform fmt -recursive
	docker compose exec terraform terraform validate
init-dev:
	docker compose exec terraform terraform -chdir=/terraform/env/dev init
plan-dev:
	docker compose exec terraform terraform -chdir=/terraform/env/dev plan
apply-dev:
	docker compose exec terraform terraform -chdir=/terraform/env/dev apply
