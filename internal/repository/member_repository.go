package repository

import (
	"context"
	"database/sql"
	"gobang/internal/domain"
)

type MemberRepository interface {
	GetMemberByUserID(ctx context.Context, userID int) (*domain.Member, error)
	CreateMember(ctx context.Context, member *domain.Member) (*int, error)
	UpdateMember(ctx context.Context, member *domain.Member) error
}

type memberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) GetMemberByUserID(ctx context.Context, userID int) (*domain.Member, error) {
	query := "SELECT id, user_id, first_name, last_name, saldo, created_at, updated_at FROM member_details WHERE user_id = $1"
	row := r.db.QueryRowContext(ctx, query, userID)
	var member domain.Member
	err := row.Scan(&member.ID, &member.UserID, &member.FirstName, &member.LastName, &member.Saldo, &member.CreatedAt, &member.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No member found with the given user ID
		}
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) CreateMember(ctx context.Context, member *domain.Member) (*int, error) {
	query := "INSERT INTO member_details (user_id, first_name, last_name, saldo) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, member.UserID, member.FirstName, member.LastName, member.Saldo).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *memberRepository) UpdateMember(ctx context.Context, member *domain.Member) error {
	query := "UPDATE member_details SET first_name = $1, last_name = $2, saldo = $3, updated_at = NOW() WHERE user_id = $4"
	_, err := r.db.ExecContext(ctx, query, member.FirstName, member.LastName, member.Saldo, member.UserID)
	return err
}