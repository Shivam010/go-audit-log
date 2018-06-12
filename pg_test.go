package logs_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Shivam010/go-audit-log"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "appointy"
	DB_NAME     = "Google"
)

func TestPostgresAuditLog(t *testing.T) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	fmt.Println("Connection Established")
	defer db.Close()
	lg, err := logs.NewPostgresAuditLog(db)
	RunAuditLogTest(lg, t)
}
