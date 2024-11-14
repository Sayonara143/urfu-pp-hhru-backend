package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ID               uuid.UUID  `json:"id"                 bun:"id,type:uuid,pk"`
	StudentProfileID uuid.UUID  `json:"student_profile_id" bun:"student_profile_id,type:uuid,notnull"`
	EmployerID       uuid.UUID  `json:"employer_id"        bun:"employer_id,type:uuid,notnull"`
	ReviewType       string     `json:"review_type"        bun:"review_type,notnull"`
	Rating           int        `json:"rating"             bun:"rating,notnull"`
	Text             string     `json:"text"               bun:"text"`
	CreatedAt        *time.Time `json:"created_at"         bun:"created_at,notnull"`
}
