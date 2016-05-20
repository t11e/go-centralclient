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

type Client interface {
	IsValidApplicationKey(key string) (bool, error)
	GetApplicationByKey(key string) (*Application, error)
	GetOrganization(id int) (*Organization, error)
	GetOrganizations() ([]*Organization, error)
	GetChildOrganizations(organization *Organization) ([]*Organization, error)
}

type client struct {
	c pc.Client
}

// New constructs a new client.
func New(pebbleClient pc.Client) (Client, error) {
	return &client{pebbleClient.Options(pc.Options{
		ServiceName: "central",
		ApiVersion:  1,
	})}, nil
}

// IsValidApplicationKey is a convenience method to check if a key is a valid
// application key. Returns true if so.
func (c *client) IsValidApplicationKey(key string) (bool, error) {
	app, err := c.GetApplicationByKey(key)
	return app != nil, err
}

// GetApplicationByKey returns an application by its key.
func (c *client) GetApplicationByKey(key string) (*Application, error) {
	var app Application
	if err := c.c.Get("/applications/keys/:key", &pc.RequestOptions{
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
func (c *client) GetOrganizations() ([]*Organization, error) {
	var organizations []*Organization
	if err := c.c.Get("/organizations", nil, &organizations); err != nil {
		return nil, err
	}
	return organizations, nil
}

// GetOrganization returns an organization by its ID. If organization does
// not exist, returns NoSuchOrganization.
func (c *client) GetOrganization(id int) (*Organization, error) {
	var organization Organization
	if err := c.c.Get("/organizations/:id", &pc.RequestOptions{
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
func (c *client) GetChildOrganizations(organization *Organization) ([]*Organization, error) {
	var organizations []*Organization
	if err := c.c.Get("/organizations/:id/organizations", &pc.RequestOptions{
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
