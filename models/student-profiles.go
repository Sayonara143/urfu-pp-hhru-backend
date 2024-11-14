package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type StudentProfile struct {
	bun.BaseModel `bun:"table:student_profiles"`

	ID         uuid.UUID  `json:"id"               bun:"id,type:uuid,pk"`
	UserID     uuid.UUID  `json:"user_id"          bun:"user_id,type:uuid,notnull"`
	Phone      string     `json:"phone"            bun:"phone"`
	University string     `json:"university"       bun:"university"`
	Faculty    string     `json:"faculty"          bun:"faculty"`
	Course     int        `json:"course"           bun:"course"`
	Skills     string     `json:"skills"           bun:"skills"`
	Languages  string     `json:"languages"        bun:"languages"`
	About      string     `json:"about"            bun:"about"`
	CreatedAt  *time.Time `json:"created_at"       bun:"created_at,notnull"`
	UpdatedAt  *time.Time `json:"updated_at"       bun:"updated_at,notnull"`

	User *User `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}
