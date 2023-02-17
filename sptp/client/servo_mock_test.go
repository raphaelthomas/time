/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: time/sptp/client/sptp.go

// Package client is a generated GoMock package.
package client

import (
	servo "github.com/facebook/time/servo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockServo is a mock of Servo interface
type MockServo struct {
	ctrl     *gomock.Controller
	recorder *MockServoMockRecorder
}

// MockServoMockRecorder is the mock recorder for MockServo
type MockServoMockRecorder struct {
	mock *MockServo
}

// NewMockServo creates a new mock instance
func NewMockServo(ctrl *gomock.Controller) *MockServo {
	mock := &MockServo{ctrl: ctrl}
	mock.recorder = &MockServoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServo) EXPECT() *MockServoMockRecorder {
	return m.recorder
}

// SyncInterval mocks base method
func (m *MockServo) SyncInterval(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SyncInterval", arg0)
}

// SyncInterval indicates an expected call of SyncInterval
func (mr *MockServoMockRecorder) SyncInterval(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncInterval", reflect.TypeOf((*MockServo)(nil).SyncInterval), arg0)
}

// Sample mocks base method
func (m *MockServo) Sample(offset int64, localTs uint64) (float64, servo.State) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sample", offset, localTs)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(servo.State)
	return ret0, ret1
}

// Sample indicates an expected call of Sample
func (mr *MockServoMockRecorder) Sample(offset, localTs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sample", reflect.TypeOf((*MockServo)(nil).Sample), offset, localTs)
}

// SetMaxFreq mocks base method
func (m *MockServo) SetMaxFreq(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxFreq", arg0)
}

// SetMaxFreq indicates an expected call of SetMaxFreq
func (mr *MockServoMockRecorder) SetMaxFreq(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxFreq", reflect.TypeOf((*MockServo)(nil).SetMaxFreq), arg0)
}