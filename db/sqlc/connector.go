package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DatabaseConnector interface {
	Connect() (*pgxpool.Pool, error)
}

type PostgresConnector struct {
	DbSource        string
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
	MaxConn         int
	MinConn         int
	ConnectTimeout  time.Duration
}

func NewPostgresConnector(dbSource string) *PostgresConnector {
	return &PostgresConnector{
		DbSource:        dbSource,
		MaxConnLifetime: 5 * time.Minute,
		MaxConnIdleTime: 5 * time.Minute,
		MaxConn:         5,
		MinConn:         1,
		ConnectTimeout:  5 * time.Second,
	}
}
func (p *PostgresConnector) Connect() (*pgxpool.Pool, error) {
	poolConn, err := pgxpool.New(context.Background(), p.DbSource)
	if err != nil {
		return nil, err
	}

	if err := poolConn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return poolConn, nil
}
