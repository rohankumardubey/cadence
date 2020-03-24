// The MIT License (MIT)
//
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package membership is a generated GoMock package.
package membership

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMonitor is a mock of Monitor interface
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorMockRecorder
}

// MockMonitorMockRecorder is the mock recorder for MockMonitor
type MockMonitorMockRecorder struct {
	mock *MockMonitor
}

// NewMockMonitor creates a new mock instance
func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &MockMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMonitor) EXPECT() *MockMonitorMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockMonitor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockMonitorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMonitor)(nil).Start))
}

// Stop mocks base method
func (m *MockMonitor) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockMonitorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMonitor)(nil).Stop))
}

// WhoAmI mocks base method
func (m *MockMonitor) WhoAmI() (*HostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhoAmI")
	ret0, _ := ret[0].(*HostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhoAmI indicates an expected call of WhoAmI
func (mr *MockMonitorMockRecorder) WhoAmI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhoAmI", reflect.TypeOf((*MockMonitor)(nil).WhoAmI))
}

// EvictSelf mocks base method
func (m *MockMonitor) EvictSelf() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvictSelf")
	ret0, _ := ret[0].(error)
	return ret0
}

// EvictSelf indicates an expected call of EvictSelf
func (mr *MockMonitorMockRecorder) EvictSelf() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictSelf", reflect.TypeOf((*MockMonitor)(nil).EvictSelf))
}

// Lookup mocks base method
func (m *MockMonitor) Lookup(service, key string) (*HostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", service, key)
	ret0, _ := ret[0].(*HostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup
func (mr *MockMonitorMockRecorder) Lookup(service, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockMonitor)(nil).Lookup), service, key)
}

// GetResolver mocks base method
func (m *MockMonitor) GetResolver(service string) (ServiceResolver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolver", service)
	ret0, _ := ret[0].(ServiceResolver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolver indicates an expected call of GetResolver
func (mr *MockMonitorMockRecorder) GetResolver(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolver", reflect.TypeOf((*MockMonitor)(nil).GetResolver), service)
}

// AddListener mocks base method
func (m *MockMonitor) AddListener(service, name string, notifyChannel chan<- *ChangedEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListener", service, name, notifyChannel)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddListener indicates an expected call of AddListener
func (mr *MockMonitorMockRecorder) AddListener(service, name, notifyChannel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListener", reflect.TypeOf((*MockMonitor)(nil).AddListener), service, name, notifyChannel)
}

// RemoveListener mocks base method
func (m *MockMonitor) RemoveListener(service, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveListener", service, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveListener indicates an expected call of RemoveListener
func (mr *MockMonitorMockRecorder) RemoveListener(service, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveListener", reflect.TypeOf((*MockMonitor)(nil).RemoveListener), service, name)
}

// GetReachableMembers mocks base method
func (m *MockMonitor) GetReachableMembers() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReachableMembers")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReachableMembers indicates an expected call of GetReachableMembers
func (mr *MockMonitorMockRecorder) GetReachableMembers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReachableMembers", reflect.TypeOf((*MockMonitor)(nil).GetReachableMembers))
}

// MockServiceResolver is a mock of ServiceResolver interface
type MockServiceResolver struct {
	ctrl     *gomock.Controller
	recorder *MockServiceResolverMockRecorder
}

// MockServiceResolverMockRecorder is the mock recorder for MockServiceResolver
type MockServiceResolverMockRecorder struct {
	mock *MockServiceResolver
}

// NewMockServiceResolver creates a new mock instance
func NewMockServiceResolver(ctrl *gomock.Controller) *MockServiceResolver {
	mock := &MockServiceResolver{ctrl: ctrl}
	mock.recorder = &MockServiceResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceResolver) EXPECT() *MockServiceResolverMockRecorder {
	return m.recorder
}

// Lookup mocks base method
func (m *MockServiceResolver) Lookup(key string) (*HostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", key)
	ret0, _ := ret[0].(*HostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup
func (mr *MockServiceResolverMockRecorder) Lookup(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockServiceResolver)(nil).Lookup), key)
}

// AddListener mocks base method
func (m *MockServiceResolver) AddListener(name string, notifyChannel chan<- *ChangedEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListener", name, notifyChannel)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddListener indicates an expected call of AddListener
func (mr *MockServiceResolverMockRecorder) AddListener(name, notifyChannel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListener", reflect.TypeOf((*MockServiceResolver)(nil).AddListener), name, notifyChannel)
}

// RemoveListener mocks base method
func (m *MockServiceResolver) RemoveListener(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveListener", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveListener indicates an expected call of RemoveListener
func (mr *MockServiceResolverMockRecorder) RemoveListener(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveListener", reflect.TypeOf((*MockServiceResolver)(nil).RemoveListener), name)
}

// MemberCount mocks base method
func (m *MockServiceResolver) MemberCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MemberCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// MemberCount indicates an expected call of MemberCount
func (mr *MockServiceResolverMockRecorder) MemberCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemberCount", reflect.TypeOf((*MockServiceResolver)(nil).MemberCount))
}

// Members mocks base method
func (m *MockServiceResolver) Members() []*HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]*HostInfo)
	return ret0
}

// Members indicates an expected call of Members
func (mr *MockServiceResolverMockRecorder) Members() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockServiceResolver)(nil).Members))
}
