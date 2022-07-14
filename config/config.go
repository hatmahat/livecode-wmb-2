package config

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	JwtSignatureKey     string
	AccessTokenLifetime time.Duration
	Client              *redis.Client
}

type GrpcConfig struct {
	Url string // set GRPC_URL=localhost:8888
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
	GrpcConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	grpcUrl := os.Getenv("GRPC_URL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.ApiConfig = ApiConfig{Url: api}

	c.TokenConfig = TokenConfig{
		ApplicationName:     "ENIGMA",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		JwtSignatureKey:     "3N!GM4",
		AccessTokenLifetime: 60 * time.Second,
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
	c.GrpcConfig = GrpcConfig{Url: grpcUrl}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
