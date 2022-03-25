package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env/v6"
)

const (
	// PostgresChannel is used by postgres trigger function by sending notifications to it
	PostgresChannel = "events"
)

// PostgresConfig is used for dialing to postgres db
type PostgresConfig struct {
	UserName string `env:"POSTGRES_USER" envDefault:"vitalijprokopenya"`
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5431"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"1234"`
	DBName   string `env:"DB_NAME" envDefault:"vitalijprokopenya"`
	URL      string
}

// GetURL returns URL to connect to postgres database
func (c *PostgresConfig) GetURL() string {
	res := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.UserName, c.Password, c.Host, c.Port, c.DBName)
	log.Info(res)
	return res
}

// GetPostgresConfig returns new config of postgresDB parsed from environment variables
func GetPostgresConfig() (*PostgresConfig, error) {
	cfg := PostgresConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error with parsing env variables in postgres config %w", err)
	}
	return &cfg, nil
}
