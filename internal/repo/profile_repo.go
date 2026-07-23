package repo

import (
	"context"

	"github.com/dimastadeoo/koda-b8-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *models.Profile) error
}

type profileRepo struct {
	pool *pgxpool.Pool
}

func NewProfileRepository(pool *pgxpool.Pool) ProfileRepository {
	return &profileRepo{pool: pool}
}

func (r *profileRepo) Create(ctx context.Context, profile *models.Profile) error {
	query := `INSERT INTO profiles (id_user, name, gender, picture, place_birth, date_birth)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`
	row := r.pool.QueryRow(ctx, query,
		profile.IDUser,
		profile.Name,
		profile.Gender,
		profile.Picture,
		profile.PlaceBirth,
		profile.DateBirth,
	)
	err := row.Scan(&profile.ID, &profile.CreatedAt, &profile.UpdatedAt)
	return err
}