package repo

import (
	"context"
	"errors"

	"github.com/dimastadeoo/koda-b8-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepository interface {
	FindByName(ctx context.Context, name string) (*models.Role, error)
}

type roleRepo struct {
	pool *pgxpool.Pool
}

func NewRoleRepository(pool *pgxpool.Pool) RoleRepository {
	return &roleRepo{pool: pool}
}

func (r *roleRepo) FindByName(ctx context.Context, name string) (*models.Role, error) {
	query := `SELECT id, name, created_at, updated_at FROM roles WHERE name = $1`
	row := r.pool.QueryRow(ctx, query, name)
	var role models.Role
	err := row.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}