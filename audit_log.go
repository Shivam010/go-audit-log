package logs

import (
	"context"
	"time"
)

// Log is an instance of an Audit Log
type Log struct {
	Timestamp time.Time
	UserID    string
	Action    string
}

// AuditLog is the interface that stores Audit Logs of actions performed by an user
type AuditLog interface {
	Add(ctx context.Context, action string) error
	GetLogsOfUser(ctx context.Context, userID string) ([]*Log, error)
	GetLogsBetweenInterval(ctx context.Context, start time.Time, end time.Time, userID string) ([]*Log, error)
}
