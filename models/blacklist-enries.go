package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BlacklistEntry struct {
	bun.BaseModel `bun:"table:blacklist_entries"`

	ID              uuid.UUID `json:"id"                 bun:"id,type:uuid,pk"`
	Type            string    `json:"type"               bun:"type,notnull"`                        // Значения: "employer_to_student", "student_to_employer"
	InitiatorUserID uuid.UUID `json:"initiator_user_id"       bun:"initiator_id,type:uuid,notnull"` // Кто заблокировал
	TargetUserID    uuid.UUID `json:"target_user_id"          bun:"target_id,type:uuid,notnull"`    // Кто был заблокирован
	Reason          string    `json:"reason"             bun:"reason"`
	CreatedAt       time.Time `json:"created_at"         bun:"created_at,notnull"`
	UpdatedAt       time.Time `json:"updated_at"         bun:"updated_at,notnull"`

	InitiatorUser *User `json:"initiator_user,omitempty" bun:"rel:belongs-to,join:initiator_id=id"`
	TargetUser    *User `json:"target_user,omitempty"    bun:"rel:belongs-to,join:target_id=id"`
}
