package postgres

import (
	"context"
	"fmt"
	"github.com/DiasOrazbaev/lrn-platform/config"
	"github.com/jackc/pgx/v5"
)

type DatabaseConfig struct {
	Db *pgx.Conn
}

func getDbUrlFromConfig(config *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DbName,
	)
}

func NewDatabase(config *config.Config) (*DatabaseConfig, error) {
	dbc := &DatabaseConfig{}

	conn, err := pgx.Connect(context.Background(), getDbUrlFromConfig(config))
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	dbc.Db = conn
	return dbc, nil
}
