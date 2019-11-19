// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira/notifier (interfaces: Notifier)

// Package mock_notifier is a generated GoMock package.
package mock_notifier

import (
	gomock "github.com/golang/mock/gomock"
	moira "github.com/moira-alert/moira"
	notifier "github.com/moira-alert/moira/notifier"
	reflect "reflect"
	sync "sync"
)

// MockNotifier is a mock of Notifier interface
type MockNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierMockRecorder
}

// MockNotifierMockRecorder is the mock recorder for MockNotifier
type MockNotifierMockRecorder struct {
	mock *MockNotifier
}

// NewMockNotifier creates a new mock instance
func NewMockNotifier(ctrl *gomock.Controller) *MockNotifier {
	mock := &MockNotifier{ctrl: ctrl}
	mock.recorder = &MockNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNotifier) EXPECT() *MockNotifierMockRecorder {
	return m.recorder
}

// GetSenders mocks base method
func (m *MockNotifier) GetSenders() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSenders")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// GetSenders indicates an expected call of GetSenders
func (mr *MockNotifierMockRecorder) GetSenders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSenders", reflect.TypeOf((*MockNotifier)(nil).GetSenders))
}

// RegisterSender mocks base method
func (m *MockNotifier) RegisterSender(arg0 map[string]string, arg1 moira.Sender) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterSender", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterSender indicates an expected call of RegisterSender
func (mr *MockNotifierMockRecorder) RegisterSender(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSender", reflect.TypeOf((*MockNotifier)(nil).RegisterSender), arg0, arg1)
}

// Send mocks base method
func (m *MockNotifier) Send(arg0 *notifier.NotificationPackage, arg1 *sync.WaitGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0, arg1)
}

// Send indicates an expected call of Send
func (mr *MockNotifierMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotifier)(nil).Send), arg0, arg1)
}

// StopSenders mocks base method
func (m *MockNotifier) StopSenders() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopSenders")
}

// StopSenders indicates an expected call of StopSenders
func (mr *MockNotifierMockRecorder) StopSenders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopSenders", reflect.TypeOf((*MockNotifier)(nil).StopSenders))
}
