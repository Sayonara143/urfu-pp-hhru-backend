package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Interview struct {
	bun.BaseModel `bun:"table:interviews"`

	ID               uuid.UUID `json:"id"                   bun:"id,type:uuid,pk"`
	JobApplicationID uuid.UUID `json:"job_application_id"   bun:"job_application_id,type:uuid,notnull"`
	StartTime        time.Time `json:"start_time"           bun:"start_time,notnull"`
	EndTime          time.Time `json:"end_time"             bun:"end_time,notnull"`
	Status           string    `json:"status"               bun:"status,notnull"`
	Location         string    `json:"location"             bun:"location"`
	Notes            string    `json:"notes"                bun:"notes"`
	CreatedAt        time.Time `json:"created_at"           bun:"created_at,notnull"`
	UpdatedAt        time.Time `json:"updated_at"           bun:"updated_at,notnull"`

	JobApplication *JobApplication `json:"job_application,omitempty" bun:"rel:belongs-to,join:job_application_id=id"`
}
