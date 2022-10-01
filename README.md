# tenhou-log

## Requirements

- golang
- docker
- nodejs
- yarn
- golang-migrate
- go-task

## Setup

```shell
make build up
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/cosmtrek/air@latest
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
# ex)arm64 macOS
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.darwin-arm64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/migrate

task migrate
make setup-terraform
```
