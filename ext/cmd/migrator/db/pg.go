package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // driver for open postgres connection

	"github.com/luanngominh/goa-example/ext/cmd/migrator/config"
)

// PGConnector store implement open for postgres
type PGConnector struct{}

// NewPGConnector create new PGConnector
func NewPGConnector() *PGConnector { return &PGConnector{} }

// Open open new connection to posgres using config
func (c *PGConnector) Open(cfg *config.Config) (*sql.DB, error) {
	sslmode := "disable"
	if cfg.DBSSLModeOption == "enable" {
		sslmode = "require"
	}

	dbstring := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s port=%s",
		cfg.DBUserName,
		cfg.DBName,
		sslmode,
		cfg.DBPassword,
		cfg.DBHostname,
		cfg.DBPort,
	)

	return sql.Open("postgres", dbstring)
}
