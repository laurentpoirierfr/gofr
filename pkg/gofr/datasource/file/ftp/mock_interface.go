// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock_interface.go -package=ftp
//

// Package ftp is a generated GoMock package.
package ftp

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLogger) Debug(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), args...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Logf mocks base method.
func (m *MockLogger) Logf(pattern string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{pattern}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockLoggerMockRecorder) Logf(pattern any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pattern}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockLogger)(nil).Logf), varargs...)
}

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// NewHistogram mocks base method.
func (m *MockMetrics) NewHistogram(name, desc string, buckets ...float64) {
	m.ctrl.T.Helper()
	varargs := []any{name, desc}
	for _, a := range buckets {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NewHistogram", varargs...)
}

// NewHistogram indicates an expected call of NewHistogram.
func (mr *MockMetricsMockRecorder) NewHistogram(name, desc any, buckets ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, desc}, buckets...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistogram", reflect.TypeOf((*MockMetrics)(nil).NewHistogram), varargs...)
}

// RecordHistogram mocks base method.
func (m *MockMetrics) RecordHistogram(ctx context.Context, name string, value float64, labels ...string) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, name, value}
	for _, a := range labels {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RecordHistogram", varargs...)
}

// RecordHistogram indicates an expected call of RecordHistogram.
func (mr *MockMetricsMockRecorder) RecordHistogram(ctx, name, value any, labels ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, name, value}, labels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordHistogram", reflect.TypeOf((*MockMetrics)(nil).RecordHistogram), varargs...)
}

// MockServerConn is a mock of ServerConn interface.
type MockServerConn struct {
	ctrl     *gomock.Controller
	recorder *MockServerConnMockRecorder
}

// MockServerConnMockRecorder is the mock recorder for MockServerConn.
type MockServerConnMockRecorder struct {
	mock *MockServerConn
}

// NewMockServerConn creates a new mock instance.
func NewMockServerConn(ctrl *gomock.Controller) *MockServerConn {
	mock := &MockServerConn{ctrl: ctrl}
	mock.recorder = &MockServerConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerConn) EXPECT() *MockServerConnMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServerConn) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServerConnMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServerConn)(nil).Delete), arg0)
}

// FileSize mocks base method.
func (m *MockServerConn) FileSize(name string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileSize", name)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileSize indicates an expected call of FileSize.
func (mr *MockServerConnMockRecorder) FileSize(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileSize", reflect.TypeOf((*MockServerConn)(nil).FileSize), name)
}

// Login mocks base method.
func (m *MockServerConn) Login(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockServerConnMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockServerConn)(nil).Login), arg0, arg1)
}

// MakeDir mocks base method.
func (m *MockServerConn) MakeDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir.
func (mr *MockServerConnMockRecorder) MakeDir(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockServerConn)(nil).MakeDir), path)
}

// Quit mocks base method.
func (m *MockServerConn) Quit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockServerConnMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockServerConn)(nil).Quit))
}

// RemoveDir mocks base method.
func (m *MockServerConn) RemoveDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDir indicates an expected call of RemoveDir.
func (mr *MockServerConnMockRecorder) RemoveDir(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDir", reflect.TypeOf((*MockServerConn)(nil).RemoveDir), path)
}

// RemoveDirRecur mocks base method.
func (m *MockServerConn) RemoveDirRecur(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDirRecur", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDirRecur indicates an expected call of RemoveDirRecur.
func (mr *MockServerConnMockRecorder) RemoveDirRecur(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDirRecur", reflect.TypeOf((*MockServerConn)(nil).RemoveDirRecur), path)
}

// Rename mocks base method.
func (m *MockServerConn) Rename(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockServerConnMockRecorder) Rename(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockServerConn)(nil).Rename), arg0, arg1)
}

// Retr mocks base method.
func (m *MockServerConn) Retr(arg0 string) (ftpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retr", arg0)
	ret0, _ := ret[0].(ftpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retr indicates an expected call of Retr.
func (mr *MockServerConnMockRecorder) Retr(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retr", reflect.TypeOf((*MockServerConn)(nil).Retr), arg0)
}

// RetrFrom mocks base method.
func (m *MockServerConn) RetrFrom(arg0 string, arg1 uint64) (ftpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrFrom", arg0, arg1)
	ret0, _ := ret[0].(ftpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrFrom indicates an expected call of RetrFrom.
func (mr *MockServerConnMockRecorder) RetrFrom(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrFrom", reflect.TypeOf((*MockServerConn)(nil).RetrFrom), arg0, arg1)
}

// Stor mocks base method.
func (m *MockServerConn) Stor(arg0 string, arg1 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stor indicates an expected call of Stor.
func (mr *MockServerConnMockRecorder) Stor(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stor", reflect.TypeOf((*MockServerConn)(nil).Stor), arg0, arg1)
}

// StorFrom mocks base method.
func (m *MockServerConn) StorFrom(arg0 string, arg1 io.Reader, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorFrom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorFrom indicates an expected call of StorFrom.
func (mr *MockServerConnMockRecorder) StorFrom(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorFrom", reflect.TypeOf((*MockServerConn)(nil).StorFrom), arg0, arg1, arg2)
}

// MockftpResponse is a mock of ftpResponse interface.
type MockftpResponse struct {
	ctrl     *gomock.Controller
	recorder *MockftpResponseMockRecorder
}

// MockftpResponseMockRecorder is the mock recorder for MockftpResponse.
type MockftpResponseMockRecorder struct {
	mock *MockftpResponse
}

// NewMockftpResponse creates a new mock instance.
func NewMockftpResponse(ctrl *gomock.Controller) *MockftpResponse {
	mock := &MockftpResponse{ctrl: ctrl}
	mock.recorder = &MockftpResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockftpResponse) EXPECT() *MockftpResponseMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockftpResponse) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockftpResponseMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockftpResponse)(nil).Close))
}

// Read mocks base method.
func (m *MockftpResponse) Read(buf []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", buf)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockftpResponseMockRecorder) Read(buf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockftpResponse)(nil).Read), buf)
}

// SetDeadline mocks base method.
func (m *MockftpResponse) SetDeadline(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline.
func (mr *MockftpResponseMockRecorder) SetDeadline(t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockftpResponse)(nil).SetDeadline), t)
}