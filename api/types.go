package api

import (
	"encoding/json"

	"github.com/google/uuid"
)

// responseWrapper used to wrap responses with pagination information
type responseWrapper struct {
	Response   any             `json:"response"`
	Pagination json.RawMessage `json:"pagination"`
}

// requestID represents struct with id.
// used as a helper struct to bind(ctx.Bind()) id from a request's path.
type requestID struct {
	ID uuid.UUID `param:"id"`
}
