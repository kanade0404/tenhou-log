package configs

import (
	"fmt"
	"os"
)

type DBEnv struct {
	Dialect  string
	Name     string
	UserName string
	Password string
	Host     string
	Port     string
	SSLMode  string
	_        struct{}
}

func (e DBEnv) ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", e.UserName, e.Password, e.Host, e.Port, e.Name, e.SSLMode)
}

type AppEnv struct {
	Port string
	_    struct{}
}

type Env struct {
	DB  DBEnv
	App AppEnv
}

func NewEnv() Env {
	return Env{
		DBEnv{
			Dialect:  os.Getenv("DB_DIALECT"),
			Name:     os.Getenv("DB_NAME"),
			UserName: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  "disable",
		},
		AppEnv{
			Port: os.Getenv("PORT"),
		},
	}
}
