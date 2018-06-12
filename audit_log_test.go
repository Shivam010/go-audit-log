package logs_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Shivam010/go-audit-log"

	"github.com/google/uuid"
)

func RunAuditLogTest(lg logs.AuditLog, t *testing.T) {
	noOfLogs := 10
	userId := uuid.New().String()
	action := "some work or action"
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userId", userId)
	var (
		strt time.Time
		end  time.Time
	)

	t.Run("Add", func(t *testing.T) {
		t.Run("One", func(t *testing.T) {
			if err := lg.Add(ctx, action); err != nil {
				t.Errorf("error while adding a log: %v", err)
			}
		})

		t.Run("Multiple Add", func(t *testing.T) {
			strt = time.Now()
			for i := 1; i <= noOfLogs; i++ {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					if err := lg.Add(ctx, action); err != nil {
						t.Errorf("error while adding a log: %v", err)
					}
				})
			}
			end = time.Now()
		})

		t.Run("Get", func(t *testing.T) {
			t.Run("Logs of userId", func(t *testing.T) {
				obj, err := lg.GetLogsOfUser(ctx, userId)
				if err != nil {
					t.Errorf("error while getting log(s) corresponding to a user %v", err)
				}

				if obj[0].UserID != userId || obj[0].Action != action {
					t.Errorf("error while getting logs of user: log retured contains wrong data")
				}

				if len(obj) != noOfLogs+1 {
					t.Errorf("error while getting logs of user: number of logs added are %v and that of got are %v", noOfLogs, len(obj))
				}
			})

			t.Run("Logs between interval", func(t *testing.T) {
				obj, err := lg.GetLogsBetweenInterval(ctx, strt, end, userId)
				if err != nil {
					t.Errorf("error while getting log(s) corresponding to a user in an interval %v", err)
				}
				// userid and action check
				if obj[0].UserID != userId || obj[0].Action != action {
					t.Errorf("error while getting logs of user: log retured contains wrong data")
				}
				// lenght of log list check
				if len(obj) != noOfLogs {
					t.Errorf("error while getting logs of user in interval: number of logs added are %v and that of got are %v", noOfLogs, len(obj))
				}
			})

			t.Run("get newUser", func(t *testing.T) {
				newUserID := uuid.New().String()
				t.Run("Logs without interval", func(t *testing.T) {
					obj, err := lg.GetLogsOfUser(ctx, newUserID)
					if err != nil {
						t.Errorf("error while getting log(s) corresponding to a user %v", err)
					}
					// length of list should be zero as no log is present for newUser
					if len(obj) != 0 {
						t.Errorf("error while getting logs of user: no data should be found but some found, %v", obj)
					}
				})
				t.Run("Logs with interval", func(t *testing.T) {
					obj, err := lg.GetLogsBetweenInterval(ctx, strt, end, newUserID)
					if err != nil {
						t.Errorf("error while getting log(s) corresponding to a user in an interval %v", err)
					}
					// length of list should be zero as no log is present for newUser
					if len(obj) != 0 {
						t.Errorf("error while getting logs of user in an interval: no data should be found but some found, %v", obj)
					}
				})
			})

			t.Run("New time Range", func(t *testing.T) {
				strt = time.Now()
				time.Sleep(5 * time.Second)
				end = time.Now()
				obj, err := lg.GetLogsBetweenInterval(ctx, strt, end, userId)
				if err != nil {
					t.Errorf("error while getting log(s) corresponding to a user in an interval %v", err)
				}
				// length of list should be zero as no log is present in time range provided
				if len(obj) != 0 {
					t.Errorf("error while getting logs of user: no data should be found but some found, %v", obj)
				}
			})
		})
	})

	t.Run("Cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		t.Run("Add", func(t *testing.T) {
			if err := lg.Add(ctx, action); err == nil {
				t.Errorf("Add should return context.Cancelled but got: %v", err)
			}
		})

		t.Run("Get logs of userId", func(t *testing.T) {
			if _, err := lg.GetLogsOfUser(ctx, userId); err == nil {
				t.Errorf("GetLogsOfUser should return context.Cancelled but got: %v", err)
			}
		})

		t.Run("Get logs in interval", func(t *testing.T) {
			if _, err := lg.GetLogsBetweenInterval(ctx, strt, end, userId); err == nil {
				t.Errorf("GetLogsBetweenInterval should return context.Cancelled but got: %v", err)
			}
		})
	})
}
