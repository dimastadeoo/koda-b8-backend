package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dimastadeoo/koda-b8-backend/internal/lib"
	"github.com/dimastadeoo/koda-b8-backend/internal/models"
	"github.com/dimastadeoo/koda-b8-backend/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, req *models.LoginRequest) (string, *models.User, error)
	Logout(ctx context.Context, token string) error
}

type authService struct {
	userRepo    repo.UserRepository
	roleRepo    repo.RoleRepository
	profileRepo repo.ProfileRepository
	sessionRepo repo.SessionRepository
}

func NewAuthService(
	userRepo repo.UserRepository,
	roleRepo repo.RoleRepository,
	profileRepo repo.ProfileRepository,
	sessionRepo repo.SessionRepository,
) AuthService {
	return &authService{
		userRepo:    userRepo,
		roleRepo:    roleRepo,
		profileRepo: profileRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *authService) Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Cari role customer
	role, err := s.roleRepo.FindByName(ctx, "customer")
	if err != nil {
		return nil, errors.New("role customer not found, please seed roles first")
	}

	// Buat user
	user := &models.User{
		IDRole:        role.ID,
		Password:      string(hashed),
		Email:         req.Email,
		HpNumber:      "", // bisa diisi nanti
		StatusAccount: "active",
	}
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Buat profile
	profile := &models.Profile{
		IDUser: user.ID,
		Name:   req.Name,
	}
	err = s.profileRepo.Create(ctx, profile)
	if err != nil {
		// Idealnya rollback user, tapi untuk simplicity kita return error
		return nil, err
	}

	// Ambil ulang user dengan relasi
	return s.userRepo.FindByID(ctx, user.ID)
}

func (s *authService) Login(ctx context.Context, req *models.LoginRequest) (string, *models.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Cek status akun
	if user.StatusAccount != "active" {
		return "", nil, fmt.Errorf("account is %s", user.StatusAccount)
	}

	// Generate JWT via lib
	tokenString, err := lib.GeneratedToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	// Simpan session
	session := &models.Session{
		IDUser: user.ID,
		Token:  tokenString,
		Status: "active",
	}
	err = s.sessionRepo.Create(ctx, session)
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func (s *authService) Logout(ctx context.Context, token string) error {
	session, err := s.sessionRepo.FindByToken(ctx, token)
	if err != nil {
		return err
	}
	return s.sessionRepo.UpdateStatus(ctx, session.ID, "inactive")
}