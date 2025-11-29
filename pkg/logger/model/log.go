package model

import "time"

// Log represents a generic log entry for tracking actions on entities
type Log struct {
	ID        int64
	Action    string
	EntityID  int64
	CreatedAt time.Time
}
