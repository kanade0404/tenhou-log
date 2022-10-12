package configs

import (
	"external/driver/secret_manager"
	"fmt"
	"os"
)

type Env struct {
	IsLocal                 bool
	AppPort                 string
	GoogleAppEnv            string
	CompressedLogBucketName string
	Dialect                 string
	DBName                  string
	UserName                string
	Password                string
	DBHost                  string
	DBPort                  string
	SSLMode                 string
	InstanceConnectionName  string
	InstancePrivateIP       string
	_                       struct{}
}

func (e *Env) ConnectionString() string {
	if e.IsLocal {
		return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", e.UserName, e.Password, e.DBHost, e.DBPort, e.DBName, e.SSLMode)
	}
	return fmt.Sprintf("user=%s password=%s host=localhost port=5432 dbname=%s sslmode=%s", e.UserName, e.Password, e.DBName, e.SSLMode)
}

func NewLocalEnv() *Env {
	return &Env{
		IsLocal:                 IsLocal(),
		AppPort:                 os.Getenv("PORT"),
		GoogleAppEnv:            os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		CompressedLogBucketName: os.Getenv("COMPRESSED_LOG_BUCKET_NAME"),
		Dialect:                 os.Getenv("DB_DIALECT"),
		DBName:                  os.Getenv("DB_NAME"),
		UserName:                os.Getenv("DB_USER"),
		Password:                os.Getenv("DB_PASSWORD"),
		DBHost:                  os.Getenv("DB_HOST"),
		DBPort:                  os.Getenv("DB_PORT"),
	}
}

func NewDevEnv(manager *secret_manager.SecretManager) (*Env, error) {
	appPort, err := manager.GetVersion("PORT")
	if err != nil {
		return nil, err
	}
	googleAppCredentials, err := manager.GetVersion("GOOGLE_APPLICATION_CREDENTIALS")
	if err != nil {
		return nil, err
	}
	compressedLogBucketName, err := manager.GetVersion("COMPRESSED_LOG_BUCKET_NAME")
	if err != nil {
		return nil, err
	}
	dialect, err := manager.GetVersion("DB_DIALECT")
	if err != nil {
		return nil, err
	}
	dbName, err := manager.GetVersion("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbUser, err := manager.GetVersion("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := manager.GetVersion("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	instanceConnName, err := manager.GetVersion("INSTANCE_CONNECTION_NAME")
	if err != nil {
		return nil, err
	}
	instancePrivateIP, err := manager.GetVersion("INSTANCE_PRIVATE_IP")
	if err != nil {
		return nil, err
	}
	return &Env{
		IsLocal:                 false,
		AppPort:                 appPort,
		GoogleAppEnv:            googleAppCredentials,
		CompressedLogBucketName: compressedLogBucketName,
		Dialect:                 dialect,
		DBName:                  dbName,
		UserName:                dbUser,
		Password:                dbPassword,
		InstanceConnectionName:  instanceConnName,
		InstancePrivateIP:       instancePrivateIP,
	}, nil
}

func IsLocal() bool {
	_, ok := os.LookupEnv("K_SERVICE")
	return !ok
}
