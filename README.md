# tenhou-log

## Requirements

- golang
- docker
- nodejs
- yarn
- sqldef
- go-task
- air
- sqlboiler

## Setup

```shell
make build up
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/cosmtrek/air@latest
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# https://github.com/k0kubun/sqldef
# ex)arm64 macOS
wget https://github.com/k0kubun/sqldef/releases/download/v0.13.12/psqldef_darwin_arm64.zip
unzip psqldef_darwin_arm64.zip
sudo mv psqldef  /usr/local/bin/sqldef

task migrate
make setup-terraform
```
