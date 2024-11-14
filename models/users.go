package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           *uuid.UUID `json:"id"             bun:"id,type:uuid,pk" param:"id"`
	Email        string     `json:"email"          bun:"email,unique,notnull"`
	FullName     string     `json:"full_name"      bun:"full_name,notnull"`
	Password     string     `json:"password,omitempty"    bun:"-"`
	PasswordHash string     `json:"-"                     bun:"password_hash"`
	PasswordSalt string     `json:"-"                     bun:"password_salt"`
	Role         string     `json:"role"           bun:"role,notnull"`
	CreatedAt    *time.Time `json:"created_at"     bun:"created_at,notnull"`
	UpdatedAt    *time.Time `json:"updated_at"     bun:"updated_at,notnull"`
}
