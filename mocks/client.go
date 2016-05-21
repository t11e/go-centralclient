package mocks

import "github.com/t11e/go-centralclient"
import "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetApplicationByKey provides a mock function with given fields: key
func (_m *Client) GetApplicationByKey(key string) (*centralclient.Application, error) {
	ret := _m.Called(key)

	var r0 *centralclient.Application
	if rf, ok := ret.Get(0).(func(string) *centralclient.Application); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*centralclient.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildOrganizations provides a mock function with given fields: organization
func (_m *Client) GetChildOrganizations(organization *centralclient.Organization) ([]*centralclient.Organization, error) {
	ret := _m.Called(organization)

	var r0 []*centralclient.Organization
	if rf, ok := ret.Get(0).(func(*centralclient.Organization) []*centralclient.Organization); ok {
		r0 = rf(organization)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*centralclient.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*centralclient.Organization) error); ok {
		r1 = rf(organization)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganization provides a mock function with given fields: id
func (_m *Client) GetOrganization(id int) (*centralclient.Organization, error) {
	ret := _m.Called(id)

	var r0 *centralclient.Organization
	if rf, ok := ret.Get(0).(func(int) *centralclient.Organization); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*centralclient.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrganizations provides a mock function with given fields:
func (_m *Client) GetOrganizations() ([]*centralclient.Organization, error) {
	ret := _m.Called()

	var r0 []*centralclient.Organization
	if rf, ok := ret.Get(0).(func() []*centralclient.Organization); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*centralclient.Organization)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsValidApplicationKey provides a mock function with given fields: key
func (_m *Client) IsValidApplicationKey(key string) (bool, error) {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
