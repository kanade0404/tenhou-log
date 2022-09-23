package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/gorilla/mux"
	"tenhou/app"
)

func main() {
	a := app.App{}
	a.Run()
}
