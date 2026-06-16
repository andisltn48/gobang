package repository

import (
	"context"
	"database/sql"
	"gobang/internal/domain"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*int, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int) error
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	query := "SELECT u.id, u.username, u.email, u.created_at, u.updated_at, m.id, m.first_name, m.last_name, m.saldo, m.created_at, m.updated_at FROM users u LEFT JOIN member_details m ON u.id = m.user_id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		var memberID sql.NullInt64
		var firstName sql.NullString
		var lastName sql.NullString
		var saldo sql.NullInt64
		var memberCreatedAt sql.NullTime
		var memberUpdatedAt sql.NullTime

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
			&memberID,
			&firstName,
			&lastName,
			&saldo,
			&memberCreatedAt,
			&memberUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Only populate member if data exists
		if memberID.Valid && memberID.Int64 != 0 {
			user.Member = &domain.Member{
				ID:        int(memberID.Int64),
				FirstName: firstName.String,
				LastName:  lastName.String,
				Saldo:     int(saldo.Int64),
				CreatedAt: memberCreatedAt.Time,
				UpdatedAt: memberUpdatedAt.Time,
			}
		}

		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	query := "SELECT u.id, u.username, u.email, u.created_at, u.updated_at, m.id, m.first_name, m.last_name, m.saldo, m.created_at, m.updated_at FROM users u LEFT JOIN member_details m ON u.id = m.user_id WHERE u.id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var user domain.User
	var memberID sql.NullInt64
	var firstName sql.NullString
	var lastName sql.NullString
	var saldo sql.NullInt64
	var memberCreatedAt sql.NullTime
	var memberUpdatedAt sql.NullTime

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&memberID,
		&firstName,
		&lastName,
		&saldo,
		&memberCreatedAt,
		&memberUpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given ID
		}
		return nil, err
	}

	// Only populate member if data exists
	if memberID.Valid && memberID.Int64 != 0 {
		user.Member = &domain.Member{
			ID:        int(memberID.Int64),
			FirstName: firstName.String,
			LastName:  lastName.String,
			Saldo:     int(saldo.Int64),
			CreatedAt: memberCreatedAt.Time,
			UpdatedAt: memberUpdatedAt.Time,
		}
	}

	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*int, error) {
	query := "INSERT INTO users (username, password, email, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"
	var userID int
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	query := "UPDATE users SET username = $1, password = $2, email = $3, updated_at = NOW() WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.ID)
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "SELECT id, username, password, email, created_at, updated_at FROM users WHERE email = $1"
	row := r.db.QueryRowContext(ctx, query, email)
	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // No user found with the given email
		}
		return nil, err
	}
	return &user, err
}
