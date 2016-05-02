package centralclient

import (
	"fmt"

	"github.com/t11e/go-pebbleclient"
)

type Client struct {
	client *pebbleclient.Client
}

// New constructs a new client.
func New(client *pebbleclient.Client) (*Client, error) {
	return &Client{client}, nil
}

// IsValidApplicationKey is a convenience method to check if a key is a valid
// application key. Returns true if so.
func (client *Client) IsValidApplicationKey(key string) (bool, error) {
	app, err := client.GetApplicationByKey(key)
	return app != nil, err
}

// GetApplicationByKey returns an application by its key.
func (client *Client) GetApplicationByKey(key string) (*Application, error) {
	var app Application
	if err := client.client.Get(fmt.Sprintf("/applications/keys/%s", key), &app); err != nil {
		if err == pebbleclient.NotFound {
			return nil, nil
		}
		return nil, err
	}
	return &app, nil
}

// GetOrganizations returns all organizations.
func (client *Client) GetOrganizations() ([]*Organization, error) {
	var organizations []*Organization
	if err := client.client.Get("/organizations", &organizations); err != nil {
		return nil, err
	}
	return organizations, nil
}

// GetOrganization returns an organization by its ID.
func (client *Client) GetOrganization(id string) (*Organization, error) {
	var organization Organization
	if err := client.client.Get(fmt.Sprintf("/organizations/%s", id), &organization); err != nil {
		if err == pebbleclient.NotFound {
			return nil, nil
		}
		return nil, err
	}
	return &organization, nil
}
