package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID                      int64  `bun:"id,pk,autoincrement"`
	Status                  string `bun:",notnull,default:pending_activation"`
	Name                    string
	Email                   string `bun:",unique,notnull"`
	PasswordHash            string `bun:",notnull"`
	ActivationCode          sql.NullString
	ActivationCodeExpiresAt sql.NullTime
	InsertedAt              time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt               time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
