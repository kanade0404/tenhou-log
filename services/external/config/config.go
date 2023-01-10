package config

import (
	"external/driver/secret_manager"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type localDatabaseConfig struct {
	dbHost string
}
type remoteDatabaseConfig struct {
	instanceConnectionName string
	instancePrivateIP      string
}
type databaseConfig struct {
	dialect  string
	dbName   string
	dbUser   string
	password string
	dbPort   string
	sslMode  string
	localDatabaseConfig
	remoteDatabaseConfig
}
type Config struct {
	isLocal                 bool
	appPort                 string
	compressedLogBucketName string
	googleAppEnv            string
	databaseConfig
}

func (c *Config) Port() string {
	return c.appPort
}

func (c *Config) Dialect() string {
	return c.dialect
}

func (c *Config) ConnectionString() string {
	if c.isLocal {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.dbUser, c.password, c.dbHost, c.dbPort, c.dbName, c.sslMode)
	}
	return fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=%s", c.dbUser, c.password, c.dbName, c.sslMode)
}
func (c *Config) CompressedLogBucketName() string {
	return c.compressedLogBucketName
}

func NewLocalEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed load env: %v", err)
	}
	return &Config{
		isLocal:                 true,
		appPort:                 os.Getenv("PORT"),
		googleAppEnv:            os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		compressedLogBucketName: os.Getenv("COMPRESSED_LOG_BUCKET_NAME"),
		databaseConfig: databaseConfig{
			dialect:  os.Getenv("DB_DIALECT"),
			dbName:   os.Getenv("DB_NAME"),
			dbPort:   os.Getenv("DB_PORT"),
			dbUser:   os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			sslMode:  "disable",
			localDatabaseConfig: localDatabaseConfig{
				dbHost: os.Getenv("DB_HOST"),
			},
		},
	}, nil
}

func NewRemoteEnv(manager *secret_manager.SecretManager) (*Config, error) {
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
	return &Config{
		isLocal:                 false,
		appPort:                 appPort,
		googleAppEnv:            googleAppCredentials,
		compressedLogBucketName: compressedLogBucketName,
		databaseConfig: databaseConfig{
			dialect:  dialect,
			dbName:   dbName,
			dbUser:   dbUser,
			password: dbPassword,
			sslMode:  "disable",
			remoteDatabaseConfig: remoteDatabaseConfig{
				instanceConnectionName: instanceConnName,
				instancePrivateIP:      instancePrivateIP,
			},
		},
	}, nil
}

func IsLocal() bool {
	_, ok := os.LookupEnv("K_SERVICE")
	return !ok
}
