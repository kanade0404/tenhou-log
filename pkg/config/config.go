package config

import (
	"context"
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/driver/secret_manager"
	"os"
)

type localDatabaseConfig struct {
	dbHost string
}
type remoteConfig struct {
	projectIDNum string
}
type databaseConfig struct {
	dialect  string
	dbName   string
	dbUser   string
	password string
	dbPort   string
	sslMode  string
	localDatabaseConfig
}
type Config struct {
	projectID               string
	isLocal                 bool
	appPort                 string
	compressedLogBucketName string
	googleAppEnv            string
	databaseConfig
	remoteConfig
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

func (c *Config) ProjectID() string {
	return c.projectID
}

func NewEnv(ctx context.Context) (*Config, error) {
	if IsLocal() {
		env, err := newLocalEnv()
		if err != nil {
			return nil, fmt.Errorf("failed to load dev env: %v", err)
		}
		return env, nil
	} else {
		sm, err := secret_manager.NewSecretManager(ctx, os.Getenv("PROJECT_ID_NUMBER"))
		if err != nil {
			return nil, fmt.Errorf("failed to initialize SecretManager: %v", err)
		}
		env, err := newRemoteEnv(sm)
		if err != nil {
			return nil, fmt.Errorf("failed load remote env: %v", err)
		}
		return env, nil
	}
}

func newLocalEnv() (*Config, error) {
	return &Config{
		isLocal:                 true,
		projectID:               os.Getenv("PROJECT_ID"),
		appPort:                 os.Getenv("SCRAPER_PORT"),
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

func newRemoteEnv(manager *secret_manager.SecretManager) (*Config, error) {
	projectID, err := manager.GetVersion("PROJECT_ID")
	if err != nil {
		return nil, err
	}
	appPort, err := manager.GetVersion("SCRAPER_PORT")
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
	return &Config{
		isLocal:                 false,
		projectID:               projectID,
		appPort:                 appPort,
		compressedLogBucketName: compressedLogBucketName,
		databaseConfig: databaseConfig{
			dialect:  dialect,
			dbName:   dbName,
			dbUser:   dbUser,
			password: dbPassword,
			sslMode:  "disable",
		},
	}, nil
}

func IsLocal() bool {
	_, ok := os.LookupEnv("K_SERVICE")
	return !ok
}
