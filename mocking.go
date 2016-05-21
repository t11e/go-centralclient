package centralclient

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

func (c *MockClient) IsValidApplicationKey(key string) (bool, error) {
	args := c.Called(key)
	return args.Bool(0), args.Error(1)
}

func (c *MockClient) GetApplicationByKey(key string) (*Application, error) {
	args := c.Called(key)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Application), nil
}

func (c *MockClient) GetOrganization(id int) (*Organization, error) {
	args := c.Called(id)
	return args.Get(0).(*Organization), args.Error(1)
}

func (c *MockClient) GetOrganizations() ([]*Organization, error) {
	args := c.Called()
	return args.Get(0).([]*Organization), args.Error(1)
}

func (c *MockClient) GetChildOrganizations(organization *Organization) ([]*Organization, error) {
	args := c.Called(organization)
	return args.Get(0).([]*Organization), args.Error(1)
}
