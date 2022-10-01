package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type LogFileEnv struct {
	CompressedLogBucketName string
}
type RuntimeEnv string

func (e RuntimeEnv) String() string {
	switch e {
	case LOCAL:
		return "local"
	case DEV:
		return "dev"
	case PROD:
		return "prod"
	default:
		return ""
	}
}

const (
	LOCAL RuntimeEnv = "local"
	DEV   RuntimeEnv = "dev"
	PROD  RuntimeEnv = "prod"
)

type AppEnv struct {
	Runtime      RuntimeEnv
	Port         string
	GoogleAppEnv string
}
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

type Env struct {
	App     AppEnv
	LogFile LogFileEnv
	DB      DBEnv
}

func NewEnv() Env {
	err := godotenv.Load(".env")
	if err != nil {
		env := os.Getenv("ENV")
		if env == "" || !(env == DEV.String() || env == PROD.String()) {
			log.Fatal(err)
		}
	}
	return Env{
		App: AppEnv{
			Port:         os.Getenv("PORT"),
			GoogleAppEnv: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		},
		LogFile: LogFileEnv{
			CompressedLogBucketName: os.Getenv("COMPRESSED_LOG_BUCKET_NAME"),
		},
		DB: DBEnv{
			Dialect:  os.Getenv("DB_DIALECT"),
			Name:     os.Getenv("DB_NAME"),
			UserName: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  "disable",
		},
	}
}
