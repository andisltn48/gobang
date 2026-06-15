package service

import (
	"context"
	"gobang/internal/domain"
	"gobang/internal/repository"
)

type MemberService interface {
	GetMemberByUserID(ctx context.Context, userID int) (*domain.Member, error)
	CreateMember(ctx context.Context, member *domain.Member) (*int, error)
	UpdateMember(ctx context.Context, member *domain.Member) error
}

type memberService struct {
	repository repository.MemberRepository
}

func NewMemberService(repository repository.MemberRepository) MemberService {
	return &memberService{repository: repository}
}

func (s *memberService) GetMemberByUserID(ctx context.Context, userID int) (*domain.Member, error) {
	return s.repository.GetMemberByUserID(ctx, userID)
}

func (s *memberService) CreateMember(ctx context.Context, member *domain.Member) (*int, error) {
	return s.repository.CreateMember(ctx, member)
}

func (s *memberService) UpdateMember(ctx context.Context, member *domain.Member) error {
	return s.repository.UpdateMember(ctx, member)
}	