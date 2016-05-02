package centralclient

import (
	"fmt"
	"net/http"
)

type ClientRequestError struct {
	Response *http.Response
}

func (err *ClientRequestError) Error() string {
	return fmt.Sprintf("Request to %s failed with status %d: %s",
		err.Response.Request.URL, err.Response.StatusCode, err.Response.Status)
}

type ClientOptions struct {
	Host string
}

type Client struct {
	host       string
	httpClient *http.Client
}

func New(options ClientOptions) (*Client, error) {
	if options.Host == "" {
		return nil, fmt.Errorf("Host must be specified")
	}
	return &Client{
		host:       options.Host,
		httpClient: &http.Client{},
	}, nil
}

func (client *Client) GetApplicationByKey(key string) (*Application, error) {
	var app Application
	if err := performJSONRequest(func() (*http.Response, error) {
		return client.httpClient.Get(client.formatUrl(fmt.Sprintf("/applications/keys/%s", key)))
	}, &app); err != nil {
		if clientErr, ok := err.(*ClientRequestError); ok && clientErr.Response.StatusCode == 404 {
			return nil, nil
		}
		return nil, err
	}
	return &app, nil
}

func (client *Client) GetOrganizations() ([]*Organization, error) {
	var organizations []*Organization
	if err := performJSONRequest(func() (*http.Response, error) {
		return client.httpClient.Get(client.formatUrl("/organizations"))
	}, &organizations); err != nil {
		return nil, err
	}
	return organizations, nil
}

func (client *Client) GetOrganization(id string) (*Organization, error) {
	var organization Organization
	if err := performJSONRequest(func() (*http.Response, error) {
		return client.httpClient.Get(client.formatUrl(fmt.Sprintf("/organizations/%s", id)))
	}, &organization); err != nil {
		if clientErr, ok := err.(*ClientRequestError); ok && clientErr.Response.StatusCode == 404 {
			return nil, nil
		}
		return nil, err
	}
	return &organization, nil
}

func (client *Client) formatUrl(path string) string {
	return fmt.Sprintf("http://%s/api/central/v1", client.host) + path
}
