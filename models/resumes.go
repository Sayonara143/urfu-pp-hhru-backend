package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Resume struct {
	bun.BaseModel `bun:"table:resumes"`

	ID               uuid.UUID `json:"id"                 bun:"id,type:uuid,pk"`
	StudentProfileID uuid.UUID `json:"student_profile_id" bun:"student_profile_id,type:uuid,notnull"`
	Title            string    `json:"title"              bun:"title,notnull"`
	Experience       string    `json:"experience"         bun:"experience"`
	Education        string    `json:"education"          bun:"education"`
	Skills           string    `json:"skills"             bun:"skills"`
	Languages        string    `json:"languages"          bun:"languages"`
	CreatedAt        time.Time `json:"created_at"         bun:"created_at,notnull"`
	UpdatedAt        time.Time `json:"updated_at"         bun:"updated_at,notnull"`

	StudentProfile *StudentProfile `json:"student_profile,omitempty" bun:"rel:belongs-to,join:student_profile_id=id"`
}
