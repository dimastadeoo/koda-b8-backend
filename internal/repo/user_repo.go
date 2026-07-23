package repo

import (
	"context"
	"time"

	"github.com/dimastadeoo/koda-b8-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
}

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepo{pool: pool}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id_role, password, email, hp_number, status_account)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	row := r.pool.QueryRow(ctx, query,
		user.IDRole,
		user.Password,
		user.Email,
		user.HpNumber,
		user.StatusAccount,
	)
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	return err
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT u.id, u.id_role, u.password, u.email, u.hp_number, u.status_account, 
		       u.created_at, u.updated_at,
		       r.id, r.name, r.created_at, r.updated_at,
		       p.id, p.id_user, p.name, p.gender, p.place_birth, p.date_birth, 
		       p.created_at, p.updated_at
		FROM users u
		LEFT JOIN roles r ON u.id_role = r.id
		LEFT JOIN profiles p ON u.id = p.id_user
		WHERE u.email = $1
	`
	row := r.pool.QueryRow(ctx, query, email)

	var user models.User
	var role models.Role
	var profile models.Profile
	var gender, placeBirth *string
	var dateBirth *time.Time

	err := row.Scan(
		&user.ID, &user.IDRole, &user.Password, &user.Email, &user.HpNumber,
		&user.StatusAccount, &user.CreatedAt, &user.UpdatedAt,
		&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt,
		&profile.ID, &profile.IDUser, &profile.Name, &gender, &placeBirth, &dateBirth,
		&profile.CreatedAt, &profile.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	profile.Gender = gender
	profile.PlaceBirth = placeBirth
	profile.DateBirth = dateBirth
	user.Role = role
	user.Profile = profile
	return &user, nil
}

func (r *userRepo) FindByID(ctx context.Context, id uint) (*models.User, error) {
	query := `
		SELECT u.id, u.id_role, u.password, u.email, u.hp_number, u.status_account,
		       u.created_at, u.updated_at,
		       r.id, r.name, r.created_at, r.updated_at,
		       p.id, p.id_user, p.name, p.gender, p.place_birth, p.date_birth,
		       p.created_at, p.updated_at
		FROM users u
		LEFT JOIN roles r ON u.id_role = r.id
		LEFT JOIN profiles p ON u.id = p.id_user
		WHERE u.id = $1
	`
	row := r.pool.QueryRow(ctx, query, id)

	var user models.User
	var role models.Role
	var profile models.Profile
	var gender, placeBirth *string
	var dateBirth *time.Time

	err := row.Scan(
		&user.ID, &user.IDRole, &user.Password, &user.Email, &user.HpNumber,
		&user.StatusAccount, &user.CreatedAt, &user.UpdatedAt,
		&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt,
		&profile.ID, &profile.IDUser, &profile.Name, &gender, &placeBirth, &dateBirth,
		&profile.CreatedAt, &profile.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	profile.Gender = gender
	profile.PlaceBirth = placeBirth
	profile.DateBirth = dateBirth
	user.Role = role
	user.Profile = profile
	return &user, nil
}