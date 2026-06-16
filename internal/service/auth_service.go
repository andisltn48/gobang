package service

import (
	"context"
	"errors"
	"gobang/internal/domain"
	"gobang/internal/dto"
	"gobang/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, request dto.RegisterDTO) (*int, error)
	Login(ctx context.Context, request dto.LoginDTO) (string, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

var jwtSecret = []byte("SUPER_SECRET_KEY_COMPANY_XYZ_2026")

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(ctx context.Context, request dto.RegisterDTO) (*int, error) {
	//check email first
	_, err := s.userRepo.FindUserByEmail(ctx, request.Email)

	if err == nil {
		return nil, errors.New("email already used")
	}

	//hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	id, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *authService) Login(ctx context.Context, request dto.LoginDTO) (string, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, request.Email)

	if err != nil {
		return "", errors.New(err.Error())
	}

	// 2. Verifikasi Bcrypt password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("wrong email or password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 Jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s *authService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
