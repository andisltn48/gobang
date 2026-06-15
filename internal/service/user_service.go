package service

import (
	"context"
	"gobang/internal/domain"
	"gobang/internal/repository"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*int, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return s.repository.GetUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repository.GetUserByID(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) (*int, error) {
	return s.repository.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.repository.UpdateUser(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	return s.repository.DeleteUser(ctx, id)
}