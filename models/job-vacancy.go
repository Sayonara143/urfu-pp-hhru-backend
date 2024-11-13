package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type JobVacancy struct {
	bun.BaseModel `bun:"table:job_vacancies"`

	ID             uuid.UUID `json:"id"               bun:"id,type:uuid,pk"`
	EmployerID     uuid.UUID `json:"employer_id"      bun:"employer_id,type:uuid,notnull"`
	Title          string    `json:"title"            bun:"title,notnull"`
	Description    string    `json:"description"      bun:"description"`
	Requirements   string    `json:"requirements"     bun:"requirements"`
	EmploymentType string    `json:"employment_type"  bun:"employment_type,notnull"`
	SalaryRange    string    `json:"salary_range"     bun:"salary_range"`
	CreatedAt      time.Time `json:"created_at"       bun:"created_at,notnull"`
	UpdatedAt      time.Time `json:"updated_at"       bun:"updated_at,notnull"`

	Employer *EmployerProfile `json:"employer,omitempty" bun:"rel:belongs-to,join:employer_id=id"`
}
