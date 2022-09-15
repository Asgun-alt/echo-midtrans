package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Provider string
	DBName   string
	Username string
	Password string
	Hostname string
	Port     string
	Timezone string
}

func NewDBConfig(provider, dbname, username, password, hostname, port, timezone string) *DBConfig {
	return &DBConfig{
		Provider: provider,
		DBName:   dbname,
		Username: username,
		Password: password,
		Hostname: hostname,
		Port:     port,
		Timezone: timezone,
	}
}

func (c *DBConfig) InitDB() (*gorm.DB, error) {
	if c.Provider == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			c.Hostname, c.Username, c.Password, c.DBName, c.Port, c.Timezone)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
		if err != nil {
			return nil, err
		}
		return db, nil
	} else {
		return nil, fmt.Errorf("unsupported RDBMS system")
	}
}
