package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SearchLog struct {
	bun.BaseModel `bun:"table:search_logs"`

	ID        uuid.UUID  `json:"id"           bun:"id,type:uuid,pk"`
	UserID    uuid.UUID  `json:"user_id"      bun:"user_id,type:uuid,notnull"`
	Query     string     `json:"query"        bun:"query,notnull"`
	CreatedAt *time.Time `json:"created_at"   bun:"created_at,notnull"`

	User *User `json:"user,omitempty" bun:"rel:belongs-to,join:user_id=id"`
}
