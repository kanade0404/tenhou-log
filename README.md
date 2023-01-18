# tenhou-log

## Requirements

- golang
- docker
- nodejs
- yarn
- ent
- go-task
- air
- direnv

## Setup

```shell
make build up
go install github.com/go-task/task/v3/cmd/task@latest
go install github.com/cosmtrek/air@latest

make setup-terraform
```
