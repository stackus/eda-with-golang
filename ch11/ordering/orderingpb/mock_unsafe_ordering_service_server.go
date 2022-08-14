// Code generated by mockery v2.14.0. DO NOT EDIT.

package orderingpb

import mock "github.com/stretchr/testify/mock"

// MockUnsafeOrderingServiceServer is an autogenerated mock type for the UnsafeOrderingServiceServer type
type MockUnsafeOrderingServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedOrderingServiceServer provides a mock function with given fields:
func (_m *MockUnsafeOrderingServiceServer) mustEmbedUnimplementedOrderingServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewMockUnsafeOrderingServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUnsafeOrderingServiceServer creates a new instance of MockUnsafeOrderingServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeOrderingServiceServer(t mockConstructorTestingTNewMockUnsafeOrderingServiceServer) *MockUnsafeOrderingServiceServer {
	mock := &MockUnsafeOrderingServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
