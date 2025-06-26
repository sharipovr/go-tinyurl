package url

import (
	"time"
)

// URL — agregat-root of domain
type URL struct {
	ID        string
	Original  string
	CreatedAt time.Time
}
