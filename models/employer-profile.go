package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EmployerProfile struct {
	bun.BaseModel `bun:"table:employer_profiles"`

	ID                 uuid.UUID `json:"id"                     bun:"id,type:uuid,pk"`
	UserID             uuid.UUID `json:"user_id"                bun:"user_id,type:uuid,notnull"`
	CompanyName        string    `json:"company_name"           bun:"company_name,notnull"`
	CompanyDescription string    `json:"company_description"    bun:"company_description"`
	Phone              string    `json:"phone"                  bun:"phone"`
	Website            string    `json:"website"                bun:"website"`
	CreatedAt          time.Time `json:"created_at"             bun:"created_at,notnull"`
	UpdatedAt          time.Time `json:"updated_at"             bun:"updated_at,notnull"`

	User *User `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}
