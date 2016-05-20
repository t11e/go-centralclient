package centralclient

import (
	"fmt"
	"net/http"

	pc "github.com/t11e/go-pebbleclient"
)

type BadAPIKey struct {
	Key string
}

func (err *BadAPIKey) Error() string {
	return fmt.Sprintf("Not a valid application key: %q", err.Key)
}

type NoSuchOrganization struct {
	Id int
}

func (err *NoSuchOrganization) Error() string {
	return fmt.Sprintf("No such organization: %q", err.Id)
}

type Client struct {
	c pc.Client
}

// New constructs a new client.
func New(client pc.Client) (*Client, error) {
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
	if err := client.c.Get("/applications/keys/:key", &pc.RequestOptions{
		Params: pc.Params{
			"key": key,
		},
	}, &app); err != nil {
		if httpErr, ok := err.(*pc.RequestError); ok && httpErr.Resp.StatusCode == http.StatusNotFound {
			return nil, &BadAPIKey{key}
		}
		return nil, err
	}
	return &app, nil
}

// GetOrganizations returns all organizations.
func (client *Client) GetOrganizations() ([]*Organization, error) {
	var organizations []*Organization
	if err := client.c.Get("/organizations", nil, &organizations); err != nil {
		return nil, err
	}
	return organizations, nil
}

// GetOrganization returns an organization by its ID. If organization does
// not exist, returns NoSuchOrganization.
func (client *Client) GetOrganization(id int) (*Organization, error) {
	var organization Organization
	if err := client.c.Get("/organizations/:id", &pc.RequestOptions{
		Params: pc.Params{
			"id": id,
		},
	}, &organization); err != nil {
		if httpErr, ok := err.(*pc.RequestError); ok && httpErr.Resp.StatusCode == http.StatusNotFound {
			return nil, &NoSuchOrganization{id}
		}
		return nil, err
	}
	return &organization, nil
}

// GetChildOrganizations returns all child organizations of another organization.
// If the organization does not exist, returns a NoSuchOrganization error.
func (client *Client) GetChildOrganizations(organization *Organization) ([]*Organization, error) {
	var organizations []*Organization
	if err := client.c.Get("/organizations/:id/organizations", &pc.RequestOptions{
		Params: pc.Params{
			"id": organization.Id,
		},
	}, &organizations); err != nil {
		if httpErr, ok := err.(*pc.RequestError); ok && httpErr.Resp.StatusCode == http.StatusNotFound {
			return nil, &NoSuchOrganization{organization.Id}
		}
		return nil, err
	}
	return organizations, nil
}
