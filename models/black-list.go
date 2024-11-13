package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Blacklist struct {
	bun.BaseModel `bun:"table:blacklist"`

	ID        uuid.UUID `json:"id"             bun:"id,type:uuid,pk"`
	UserID    uuid.UUID `json:"user_id"        bun:"user_id,type:uuid,notnull"`
	Reason    string    `json:"reason"         bun:"reason"`
	BannedBy  uuid.UUID `json:"banned_by"      bun:"banned_by,type:uuid,notnull"`
	Permanent bool      `json:"permanent"      bun:"permanent,notnull"`
	CreatedAt time.Time `json:"created_at"     bun:"created_at,notnull"`
	UpdatedAt time.Time `json:"updated_at"     bun:"updated_at,notnull"`

	User         *User `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
	BannedByUser *User `json:"banned_by_user,omitempty" bun:"rel:belongs-to,join:banned_by=id"`
}
