package logs

import (
	"context"
	"database/sql"
	"time"
)

type pg struct {
	*sql.DB
}

// setupDatabase sets the PostgreSQL database to use Timescale DB
func setupDatabase(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Creating a timescaledb extension for the database
	const ext = `CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;`
	if _, err = tx.Exec(ext); err != nil {
		return err
	}

	// creating the audit log table
	const tbl = `CREATE TABLE IF NOT EXISTS audit."Logs" (
		"Timestamp"        TIMESTAMPTZ       NOT NULL,
		"UserId" text NOT NULL,
		"ServiceName" text NOT NULL,
		"Action" text NOT NULL
	  );`
	if _, err = tx.Exec(tbl); err != nil {
		return err
	}

	// creating the hypertable of audit log table for timescaledb
	const hptbl = `SELECT create_hypertable('audit."Logs"', 'Timestamp',if_not_exists => true);`
	if _, err = tx.Exec(hptbl); err != nil {
		return err
	}
	return nil
}

// NewPostgresAuditLog returns the AuditLog interface that implements Audit Logs Functions
func NewPostgresAuditLog(db *sql.DB) (AuditLog, error) {
	if err := setupDatabase(db); err != nil {
		return nil, err
	}
	return &pg{db}, nil
}

// Add adds a log into database
func (db *pg) Add(ctx context.Context, action string) error {
	userID := ctx.Value("userId").(string)
	const stmt = `INSERT INTO audit."Logs" VALUES (NOW(), $1, $2);`
	if _, err := db.ExecContext(ctx, stmt, userID, action); err != nil {
		return err
	}
	return nil
}

// GetLogsOfUser returns the list of all the Logs for a user
func (db *pg) GetLogsOfUser(ctx context.Context, userID string) ([]*Log, error) {
	const stmt = `SELECT * FROM audit."Logs" WHERE "UserId" = $1;`
	rows, err := db.QueryContext(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}
	lst := make([]*Log, 0, 100)
	for rows.Next() {
		l := Log{}
		if err := rows.Scan(&l.Timestamp, &l.UserID, &l.Action); err != nil {
			return nil, err
		}
		lst = append(lst, &l)
	}
	return lst, nil
}

// GetLogsBetweenInterval returns the list of all the Logs for a user in a range of interval
func (db *pg) GetLogsBetweenInterval(ctx context.Context, start time.Time, end time.Time, userID string) ([]*Log, error) {
	const stmt = `SELECT * FROM audit."Logs" WHERE "UserId" = $1 AND "Timestamp" >= $2 AND "Timestamp" <= $3;`
	rows, err := db.QueryContext(ctx, stmt, userID, start, end)
	if err != nil {
		return nil, err
	}
	lst := make([]*Log, 0, 100)
	for rows.Next() {
		l := Log{}
		if err := rows.Scan(&l.Timestamp, &l.UserID, &l.Action); err != nil {
			return nil, err
		}
		lst = append(lst, &l)
	}
	return lst, nil
}
