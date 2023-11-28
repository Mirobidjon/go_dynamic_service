package postgres

import (
	"context"
	pd "kassa360/kassa360_go_dynamic_service/genproto/dynamic_service"
	"kassa360/kassa360_go_dynamic_service/storage"

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
