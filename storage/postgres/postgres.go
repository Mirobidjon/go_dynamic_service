package postgres

import (
	"context"
	"fmt"
	"kassa360/kassa360_go_dynamic_service/config"
	"kassa360/kassa360_go_dynamic_service/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storagePg struct {
	db         *pgxpool.Pool
	entityRepo storage.EntityPgI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StoragePgI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	config.ConnConfig.RuntimeParams = map[string]string{
		"application_name":            cfg.ServiceName,
		"statement_timeout":           "30000",
		"standard_conforming_strings": "on",
		"TimeZone":                    "UTC",
	}
	// config.ConnConfig.PreferSimpleProtocol = true

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &storagePg{
		db: pool,
	}, err
}

func (s *storagePg) CloseDB() {
	s.db.Close()
}

func (s *storagePg) Entity() storage.EntityPgI {
	if s.entityRepo == nil {
		s.entityRepo = NewEntityRepo(s.db)
	}

	return s.entityRepo
}
