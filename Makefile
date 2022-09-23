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
down-tenhou-scraper:
	docker compose down tenhou_scraper
down-log-db:
	docker compose down tenhou_log_db
down-all:
	docker-compose down --rmi local --volumes --remove-orphans
exec-tenhou-scraper:
	docker compose exec tenhou_scraper bash
exec-terraform:
	docker compose exec terraform bash
log-tenhou-scraper:
	docker compose logs tenhou_scraper
log-db:
	docker compose logs tenhou_log_db
setup-terraform:
	docker compose exec terraform gcloud auth application-default login
	docker compose exec terraform gcloud auth login
	make project-set
	docker compose exec terraform terraform init
project-set:
	docker compose exec terraform gcloud config set project kanade0404
format:
	docker compose exec terraform terraform fmt -recursive
	docker compose exec terraform terraform validate
