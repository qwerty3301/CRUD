package repository

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jackc/pgx/v5"
	"log"
)

type Config struct {
	Database DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	Password string
	Username string `toml:"Username"`
	Port     string `toml:"Port"`
	Database string `toml:"Database"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (p *DatabaseConfig) NewPostgreDB(config *DatabaseConfig) *DatabaseConfig {
	p.Password = config.Password
	p.Username = config.Username
	p.Port = config.Port
	p.Database = config.Database
	return p
}

func (p *DatabaseConfig) Connect() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@db:%s/%s",
		p.Username, p.Password, p.Port, p.Database))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return conn
}
