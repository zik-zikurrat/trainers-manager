package postgres

import (
	"context"
	"trainers-manager/internal/config"
	"trainers-manager/internal/helper"
	"trainers-manager/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres -.
type Posgtres struct {
	log  *logger.Logger
	Pool *pgxpool.Pool
}

// New -.
func New(cfg *config.Config, log *logger.Logger) (*Posgtres, error) {
	pool, err := pgxpool.New(context.Background(), helper.GetDBDsn(cfg))
	if err != nil {
		log.Error("failed to create pg pool", err)
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		log.Error("failed to ping db", err)
		return nil, err
	}
	log.Info("database connected successfully")
	return &Posgtres{
		log:  log,
		Pool: pool,
	}, nil
}

// Close -.
func (p *Posgtres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
