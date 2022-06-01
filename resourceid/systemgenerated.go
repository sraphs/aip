package resourceid

import (
	"github.com/google/uuid"
)

// Next returns a new system-generated resource ID.
func Next() string {
	return uuid.New().String()
}
