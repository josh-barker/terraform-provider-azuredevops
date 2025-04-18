// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/servicehooks (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	forminput "github.com/microsoft/azure-devops-go-api/azuredevops/v7/forminput"
	notification "github.com/microsoft/azure-devops-go-api/azuredevops/v7/notification"
	servicehooks "github.com/microsoft/azure-devops-go-api/azuredevops/v7/servicehooks"
	gomock "go.uber.org/mock/gomock"
)

// MockServicehooksClient is a mock of Client interface.
type MockServicehooksClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicehooksClientMockRecorder
	isgomock struct{}
}

// MockServicehooksClientMockRecorder is the mock recorder for MockServicehooksClient.
type MockServicehooksClientMockRecorder struct {
	mock *MockServicehooksClient
}

// NewMockServicehooksClient creates a new mock instance.
func NewMockServicehooksClient(ctrl *gomock.Controller) *MockServicehooksClient {
	mock := &MockServicehooksClient{ctrl: ctrl}
	mock.recorder = &MockServicehooksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicehooksClient) EXPECT() *MockServicehooksClientMockRecorder {
	return m.recorder
}

// CreateSubscription mocks base method.
func (m *MockServicehooksClient) CreateSubscription(arg0 context.Context, arg1 servicehooks.CreateSubscriptionArgs) (*servicehooks.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscription", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscription indicates an expected call of CreateSubscription.
func (mr *MockServicehooksClientMockRecorder) CreateSubscription(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscription", reflect.TypeOf((*MockServicehooksClient)(nil).CreateSubscription), arg0, arg1)
}

// CreateSubscriptionsQuery mocks base method.
func (m *MockServicehooksClient) CreateSubscriptionsQuery(arg0 context.Context, arg1 servicehooks.CreateSubscriptionsQueryArgs) (*servicehooks.SubscriptionsQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscriptionsQuery", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.SubscriptionsQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscriptionsQuery indicates an expected call of CreateSubscriptionsQuery.
func (mr *MockServicehooksClientMockRecorder) CreateSubscriptionsQuery(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscriptionsQuery", reflect.TypeOf((*MockServicehooksClient)(nil).CreateSubscriptionsQuery), arg0, arg1)
}

// CreateTestNotification mocks base method.
func (m *MockServicehooksClient) CreateTestNotification(arg0 context.Context, arg1 servicehooks.CreateTestNotificationArgs) (*servicehooks.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestNotification", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTestNotification indicates an expected call of CreateTestNotification.
func (mr *MockServicehooksClientMockRecorder) CreateTestNotification(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestNotification", reflect.TypeOf((*MockServicehooksClient)(nil).CreateTestNotification), arg0, arg1)
}

// DeleteSubscription mocks base method.
func (m *MockServicehooksClient) DeleteSubscription(arg0 context.Context, arg1 servicehooks.DeleteSubscriptionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription.
func (mr *MockServicehooksClientMockRecorder) DeleteSubscription(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscription", reflect.TypeOf((*MockServicehooksClient)(nil).DeleteSubscription), arg0, arg1)
}

// GetConsumer mocks base method.
func (m *MockServicehooksClient) GetConsumer(arg0 context.Context, arg1 servicehooks.GetConsumerArgs) (*servicehooks.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumer", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumer indicates an expected call of GetConsumer.
func (mr *MockServicehooksClientMockRecorder) GetConsumer(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumer", reflect.TypeOf((*MockServicehooksClient)(nil).GetConsumer), arg0, arg1)
}

// GetConsumerAction mocks base method.
func (m *MockServicehooksClient) GetConsumerAction(arg0 context.Context, arg1 servicehooks.GetConsumerActionArgs) (*servicehooks.ConsumerAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumerAction", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.ConsumerAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumerAction indicates an expected call of GetConsumerAction.
func (mr *MockServicehooksClientMockRecorder) GetConsumerAction(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumerAction", reflect.TypeOf((*MockServicehooksClient)(nil).GetConsumerAction), arg0, arg1)
}

// GetEventType mocks base method.
func (m *MockServicehooksClient) GetEventType(arg0 context.Context, arg1 servicehooks.GetEventTypeArgs) (*servicehooks.EventTypeDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventType", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.EventTypeDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventType indicates an expected call of GetEventType.
func (mr *MockServicehooksClientMockRecorder) GetEventType(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventType", reflect.TypeOf((*MockServicehooksClient)(nil).GetEventType), arg0, arg1)
}

// GetNotification mocks base method.
func (m *MockServicehooksClient) GetNotification(arg0 context.Context, arg1 servicehooks.GetNotificationArgs) (*servicehooks.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotification", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotification indicates an expected call of GetNotification.
func (mr *MockServicehooksClientMockRecorder) GetNotification(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotification", reflect.TypeOf((*MockServicehooksClient)(nil).GetNotification), arg0, arg1)
}

// GetNotifications mocks base method.
func (m *MockServicehooksClient) GetNotifications(arg0 context.Context, arg1 servicehooks.GetNotificationsArgs) (*[]servicehooks.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", arg0, arg1)
	ret0, _ := ret[0].(*[]servicehooks.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockServicehooksClientMockRecorder) GetNotifications(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockServicehooksClient)(nil).GetNotifications), arg0, arg1)
}

// GetPublisher mocks base method.
func (m *MockServicehooksClient) GetPublisher(arg0 context.Context, arg1 servicehooks.GetPublisherArgs) (*servicehooks.Publisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublisher", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublisher indicates an expected call of GetPublisher.
func (mr *MockServicehooksClientMockRecorder) GetPublisher(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublisher", reflect.TypeOf((*MockServicehooksClient)(nil).GetPublisher), arg0, arg1)
}

// GetSubscription mocks base method.
func (m *MockServicehooksClient) GetSubscription(arg0 context.Context, arg1 servicehooks.GetSubscriptionArgs) (*servicehooks.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscription", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscription indicates an expected call of GetSubscription.
func (mr *MockServicehooksClientMockRecorder) GetSubscription(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscription", reflect.TypeOf((*MockServicehooksClient)(nil).GetSubscription), arg0, arg1)
}

// GetSubscriptionDiagnostics mocks base method.
func (m *MockServicehooksClient) GetSubscriptionDiagnostics(arg0 context.Context, arg1 servicehooks.GetSubscriptionDiagnosticsArgs) (*notification.SubscriptionDiagnostics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionDiagnostics", arg0, arg1)
	ret0, _ := ret[0].(*notification.SubscriptionDiagnostics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionDiagnostics indicates an expected call of GetSubscriptionDiagnostics.
func (mr *MockServicehooksClientMockRecorder) GetSubscriptionDiagnostics(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionDiagnostics", reflect.TypeOf((*MockServicehooksClient)(nil).GetSubscriptionDiagnostics), arg0, arg1)
}

// ListConsumerActions mocks base method.
func (m *MockServicehooksClient) ListConsumerActions(arg0 context.Context, arg1 servicehooks.ListConsumerActionsArgs) (*[]servicehooks.ConsumerAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConsumerActions", arg0, arg1)
	ret0, _ := ret[0].(*[]servicehooks.ConsumerAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConsumerActions indicates an expected call of ListConsumerActions.
func (mr *MockServicehooksClientMockRecorder) ListConsumerActions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConsumerActions", reflect.TypeOf((*MockServicehooksClient)(nil).ListConsumerActions), arg0, arg1)
}

// ListConsumers mocks base method.
func (m *MockServicehooksClient) ListConsumers(arg0 context.Context, arg1 servicehooks.ListConsumersArgs) (*[]servicehooks.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConsumers", arg0, arg1)
	ret0, _ := ret[0].(*[]servicehooks.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConsumers indicates an expected call of ListConsumers.
func (mr *MockServicehooksClientMockRecorder) ListConsumers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConsumers", reflect.TypeOf((*MockServicehooksClient)(nil).ListConsumers), arg0, arg1)
}

// ListEventTypes mocks base method.
func (m *MockServicehooksClient) ListEventTypes(arg0 context.Context, arg1 servicehooks.ListEventTypesArgs) (*[]servicehooks.EventTypeDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEventTypes", arg0, arg1)
	ret0, _ := ret[0].(*[]servicehooks.EventTypeDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEventTypes indicates an expected call of ListEventTypes.
func (mr *MockServicehooksClientMockRecorder) ListEventTypes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventTypes", reflect.TypeOf((*MockServicehooksClient)(nil).ListEventTypes), arg0, arg1)
}

// ListPublishers mocks base method.
func (m *MockServicehooksClient) ListPublishers(arg0 context.Context, arg1 servicehooks.ListPublishersArgs) (*[]servicehooks.Publisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublishers", arg0, arg1)
	ret0, _ := ret[0].(*[]servicehooks.Publisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublishers indicates an expected call of ListPublishers.
func (mr *MockServicehooksClientMockRecorder) ListPublishers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublishers", reflect.TypeOf((*MockServicehooksClient)(nil).ListPublishers), arg0, arg1)
}

// ListSubscriptions mocks base method.
func (m *MockServicehooksClient) ListSubscriptions(arg0 context.Context, arg1 servicehooks.ListSubscriptionsArgs) (*[]servicehooks.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubscriptions", arg0, arg1)
	ret0, _ := ret[0].(*[]servicehooks.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubscriptions indicates an expected call of ListSubscriptions.
func (mr *MockServicehooksClientMockRecorder) ListSubscriptions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscriptions", reflect.TypeOf((*MockServicehooksClient)(nil).ListSubscriptions), arg0, arg1)
}

// QueryInputValues mocks base method.
func (m *MockServicehooksClient) QueryInputValues(arg0 context.Context, arg1 servicehooks.QueryInputValuesArgs) (*forminput.InputValuesQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryInputValues", arg0, arg1)
	ret0, _ := ret[0].(*forminput.InputValuesQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryInputValues indicates an expected call of QueryInputValues.
func (mr *MockServicehooksClientMockRecorder) QueryInputValues(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryInputValues", reflect.TypeOf((*MockServicehooksClient)(nil).QueryInputValues), arg0, arg1)
}

// QueryNotifications mocks base method.
func (m *MockServicehooksClient) QueryNotifications(arg0 context.Context, arg1 servicehooks.QueryNotificationsArgs) (*servicehooks.NotificationsQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryNotifications", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.NotificationsQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryNotifications indicates an expected call of QueryNotifications.
func (mr *MockServicehooksClientMockRecorder) QueryNotifications(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryNotifications", reflect.TypeOf((*MockServicehooksClient)(nil).QueryNotifications), arg0, arg1)
}

// QueryPublishers mocks base method.
func (m *MockServicehooksClient) QueryPublishers(arg0 context.Context, arg1 servicehooks.QueryPublishersArgs) (*servicehooks.PublishersQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryPublishers", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.PublishersQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryPublishers indicates an expected call of QueryPublishers.
func (mr *MockServicehooksClientMockRecorder) QueryPublishers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryPublishers", reflect.TypeOf((*MockServicehooksClient)(nil).QueryPublishers), arg0, arg1)
}

// ReplaceSubscription mocks base method.
func (m *MockServicehooksClient) ReplaceSubscription(arg0 context.Context, arg1 servicehooks.ReplaceSubscriptionArgs) (*servicehooks.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceSubscription", arg0, arg1)
	ret0, _ := ret[0].(*servicehooks.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceSubscription indicates an expected call of ReplaceSubscription.
func (mr *MockServicehooksClientMockRecorder) ReplaceSubscription(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceSubscription", reflect.TypeOf((*MockServicehooksClient)(nil).ReplaceSubscription), arg0, arg1)
}

// UpdateSubscriptionDiagnostics mocks base method.
func (m *MockServicehooksClient) UpdateSubscriptionDiagnostics(arg0 context.Context, arg1 servicehooks.UpdateSubscriptionDiagnosticsArgs) (*notification.SubscriptionDiagnostics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriptionDiagnostics", arg0, arg1)
	ret0, _ := ret[0].(*notification.SubscriptionDiagnostics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubscriptionDiagnostics indicates an expected call of UpdateSubscriptionDiagnostics.
func (mr *MockServicehooksClientMockRecorder) UpdateSubscriptionDiagnostics(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriptionDiagnostics", reflect.TypeOf((*MockServicehooksClient)(nil).UpdateSubscriptionDiagnostics), arg0, arg1)
}
