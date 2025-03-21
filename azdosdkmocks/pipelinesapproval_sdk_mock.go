// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelinesapproval (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	pipelinesapproval "github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelinesapproval"
	gomock "go.uber.org/mock/gomock"
)

// MockPipelinesapprovalClient is a mock of Client interface.
type MockPipelinesapprovalClient struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesapprovalClientMockRecorder
	isgomock struct{}
}

// MockPipelinesapprovalClientMockRecorder is the mock recorder for MockPipelinesapprovalClient.
type MockPipelinesapprovalClientMockRecorder struct {
	mock *MockPipelinesapprovalClient
}

// NewMockPipelinesapprovalClient creates a new mock instance.
func NewMockPipelinesapprovalClient(ctrl *gomock.Controller) *MockPipelinesapprovalClient {
	mock := &MockPipelinesapprovalClient{ctrl: ctrl}
	mock.recorder = &MockPipelinesapprovalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesapprovalClient) EXPECT() *MockPipelinesapprovalClientMockRecorder {
	return m.recorder
}

// GetApproval mocks base method.
func (m *MockPipelinesapprovalClient) GetApproval(arg0 context.Context, arg1 pipelinesapproval.GetApprovalArgs) (*pipelinesapproval.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApproval", arg0, arg1)
	ret0, _ := ret[0].(*pipelinesapproval.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApproval indicates an expected call of GetApproval.
func (mr *MockPipelinesapprovalClientMockRecorder) GetApproval(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApproval", reflect.TypeOf((*MockPipelinesapprovalClient)(nil).GetApproval), arg0, arg1)
}

// QueryApprovals mocks base method.
func (m *MockPipelinesapprovalClient) QueryApprovals(arg0 context.Context, arg1 pipelinesapproval.QueryApprovalsArgs) (*[]pipelinesapproval.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryApprovals", arg0, arg1)
	ret0, _ := ret[0].(*[]pipelinesapproval.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryApprovals indicates an expected call of QueryApprovals.
func (mr *MockPipelinesapprovalClientMockRecorder) QueryApprovals(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryApprovals", reflect.TypeOf((*MockPipelinesapprovalClient)(nil).QueryApprovals), arg0, arg1)
}

// UpdateApprovals mocks base method.
func (m *MockPipelinesapprovalClient) UpdateApprovals(arg0 context.Context, arg1 pipelinesapproval.UpdateApprovalsArgs) (*[]pipelinesapproval.Approval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApprovals", arg0, arg1)
	ret0, _ := ret[0].(*[]pipelinesapproval.Approval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApprovals indicates an expected call of UpdateApprovals.
func (mr *MockPipelinesapprovalClientMockRecorder) UpdateApprovals(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApprovals", reflect.TypeOf((*MockPipelinesapprovalClient)(nil).UpdateApprovals), arg0, arg1)
}
