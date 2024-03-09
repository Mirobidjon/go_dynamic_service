package postgres

import (
	"context"

	pd "github.com/mirobidjon/go_dynamic_service/genproto/dynamic_service"
	"github.com/mirobidjon/go_dynamic_service/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type entityRepo struct {
	db *pgxpool.Pool
}

func NewEntityRepo(db *pgxpool.Pool) storage.EntityPgI {
	return &entityRepo{
		db: db,
	}
}

func (r *entityRepo) Create(ctx context.Context, group *pd.Group, body map[string]interface{}) error {
	return nil
}
