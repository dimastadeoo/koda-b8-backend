package repo

import (
	"context"

	"github.com/dimastadeoo/koda-b8-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepository interface {
	Create(ctx context.Context, session *models.Session) error
	FindByToken(ctx context.Context, token string) (*models.Session, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
	DeleteByToken(ctx context.Context, token string) error
}

type sessionRepo struct {
	pool *pgxpool.Pool
}

func NewSessionRepository(pool *pgxpool.Pool) SessionRepository {
	return &sessionRepo{pool: pool}
}

func (r *sessionRepo) Create(ctx context.Context, session *models.Session) error {
	query := `INSERT INTO sessions (id_user, token, status) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	row := r.pool.QueryRow(ctx, query, session.IDUser, session.Token, session.Status)
	return row.Scan(&session.ID, &session.CreatedAt, &session.UpdatedAt)
}

func (r *sessionRepo) FindByToken(ctx context.Context, token string) (*models.Session, error) {
	query := `SELECT id, id_user, token, status, created_at, updated_at FROM sessions WHERE token = $1`
	row := r.pool.QueryRow(ctx, query, token)
	var session models.Session
	err := row.Scan(&session.ID, &session.IDUser, &session.Token, &session.Status, &session.CreatedAt, &session.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepo) UpdateStatus(ctx context.Context, id uint, status string) error {
	query := `UPDATE sessions SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.pool.Exec(ctx, query, status, id)
	return err
}

func (r *sessionRepo) DeleteByToken(ctx context.Context, token string) error {
	query := `DELETE FROM sessions WHERE token = $1`
	_, err := r.pool.Exec(ctx, query, token)
	return err
}