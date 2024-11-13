package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type JobApplication struct {
	bun.BaseModel `bun:"table:job_applications"`

	ID               uuid.UUID `json:"id"                 bun:"id,type:uuid,pk"`
	JobVacancyID     uuid.UUID `json:"job_vacancy_id"     bun:"job_vacancy_id,type:uuid,notnull"`
	StudentProfileID uuid.UUID `json:"student_profile_id" bun:"student_profile_id,type:uuid,notnull"`
	CoverLetter      string    `json:"cover_letter"       bun:"cover_letter"`
	Status           string    `json:"status"             bun:"status,notnull"`
	CreatedAt        time.Time `json:"created_at"         bun:"created_at,notnull"`
	UpdatedAt        time.Time `json:"updated_at"         bun:"updated_at,notnull"`

	JobVacancy     *JobVacancy     `json:"job_vacancy,omitempty" bun:"rel:belongs-to,join:job_vacancy_id=id"`
	StudentProfile *StudentProfile `json:"student_profile,omitempty" bun:"rel:belongs-to,join:student_profile_id=id"`
}
