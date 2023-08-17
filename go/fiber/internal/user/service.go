package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ygunayer/restbench/internal/database"
	"github.com/ygunayer/restbench/internal/passwordhasher"
)

const ActivationCodeExpiryPeriod = 3 * 24 * time.Hour

func RegisterUser(ctx context.Context, c RegisterUserCommand) (*User, error) {
	activationCode := uuid.New().String()
	passwordHash := passwordhasher.HashPassword(c.Password)

	now := time.Now()

	user := User{
		Name:                    c.Name,
		Email:                   c.Email,
		ActivationCode:          sql.NullString{Valid: true, String: activationCode},
		ActivationCodeExpiresAt: sql.NullTime{Valid: true, Time: now.Add(ActivationCodeExpiryPeriod)},
		InsertedAt:              now,
		UpdatedAt:               now,
		PasswordHash:            passwordHash,
	}

	_, err := database.Get().NewInsert().Model(&user).Returning("*").Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
