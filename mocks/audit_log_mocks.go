// Code generated by MockGen. DO NOT EDIT.
// Source: audit_log.go

// Package logs_mock is a generated GoMock package.
package logs_mock

import (
	context "context"
	go_audit_log "github.com/Shivam010/go-audit-log"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockAuditLog is a mock of AuditLog interface
type MockAuditLog struct {
	ctrl     *gomock.Controller
	recorder *MockAuditLogMockRecorder
}

// MockAuditLogMockRecorder is the mock recorder for MockAuditLog
type MockAuditLogMockRecorder struct {
	mock *MockAuditLog
}

// NewMockAuditLog creates a new mock instance
func NewMockAuditLog(ctrl *gomock.Controller) *MockAuditLog {
	mock := &MockAuditLog{ctrl: ctrl}
	mock.recorder = &MockAuditLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuditLog) EXPECT() *MockAuditLogMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockAuditLog) Add(ctx context.Context, action string) error {
	ret := m.ctrl.Call(m, "Add", ctx, action)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockAuditLogMockRecorder) Add(ctx, action interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAuditLog)(nil).Add), ctx, action)
}

// GetLogsOfUser mocks base method
func (m *MockAuditLog) GetLogsOfUser(ctx context.Context, userID string) ([]*go_audit_log.Log, error) {
	ret := m.ctrl.Call(m, "GetLogsOfUser", ctx, userID)
	ret0, _ := ret[0].([]*go_audit_log.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogsOfUser indicates an expected call of GetLogsOfUser
func (mr *MockAuditLogMockRecorder) GetLogsOfUser(ctx, userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsOfUser", reflect.TypeOf((*MockAuditLog)(nil).GetLogsOfUser), ctx, userID)
}

// GetLogsBetweenInterval mocks base method
func (m *MockAuditLog) GetLogsBetweenInterval(ctx context.Context, start, end time.Time, userID string) ([]*go_audit_log.Log, error) {
	ret := m.ctrl.Call(m, "GetLogsBetweenInterval", ctx, start, end, userID)
	ret0, _ := ret[0].([]*go_audit_log.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogsBetweenInterval indicates an expected call of GetLogsBetweenInterval
func (mr *MockAuditLogMockRecorder) GetLogsBetweenInterval(ctx, start, end, userID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsBetweenInterval", reflect.TypeOf((*MockAuditLog)(nil).GetLogsBetweenInterval), ctx, start, end, userID)
}
