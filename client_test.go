package centralclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t11e/go-pebbleclient"
)

func newTestClient(host string) (*Client, error) {
	pebClient, err := pebbleclient.New(pebbleclient.ClientOptions{
		Host:    host,
		AppName: "central",
	})
	if err != nil {
		return nil, err
	}

	return New(pebClient)
}

func TestNew_valid(t *testing.T) {
	client, err := newTestClient("localhost")
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestClient_GetApplicationByKey_found(t *testing.T) {
	org := &Application{
		Id:   1,
		Name: "The Illuminati & Co.",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/api/central/v1/applications/keys/frobnitz", req.URL.Path)
		encoder := json.NewEncoder(w)
		encoder.Encode(org)
	}))

	client, err := newTestClient(hostFromUrl(server.URL))
	assert.NoError(t, err)

	resultApp, err := client.GetApplicationByKey("frobnitz")
	assert.NoError(t, err)
	assert.Equal(t, org, resultApp)
}

func TestClient_GetApplicationByKey_notFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(404)
	}))

	client, err := newTestClient(hostFromUrl(server.URL))
	assert.NoError(t, err)

	org, err := client.GetApplicationByKey("frobnitz")
	assert.NoError(t, err)
	assert.Nil(t, org)
}

func TestClient_GetOrganization_found(t *testing.T) {
	org := &Organization{
		Id:    1,
		Title: "The Illuminati",
		Realm: "illuminati",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/api/central/v1/organizations/1", req.URL.Path)
		encoder := json.NewEncoder(w)
		encoder.Encode(org)
	}))

	client, err := newTestClient(hostFromUrl(server.URL))
	assert.NoError(t, err)

	resultOrg, err := client.GetOrganization(1)
	assert.NoError(t, err)
	assert.Equal(t, org, resultOrg)
}

func TestClient_GetOrganization_notFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(404)
	}))

	client, err := newTestClient(hostFromUrl(server.URL))
	assert.NoError(t, err)

	org, err := client.GetOrganization(1)
	assert.NoError(t, err)
	assert.Nil(t, org)
}

func TestClient_GetOrganizations(t *testing.T) {
	orgs := []*Organization{
		&Organization{
			Id:    1,
			Title: "The Illuminati",
			Realm: "illuminati",
		},
		&Organization{
			Id:    2,
			Title: "NSA",
			Realm: "nsa",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/api/central/v1/organizations", req.URL.Path)
		encoder := json.NewEncoder(w)
		encoder.Encode(orgs)
	}))

	client, err := newTestClient(hostFromUrl(server.URL))
	assert.NoError(t, err)

	resultOrgs, err := client.GetOrganizations()
	assert.NoError(t, err)
	assert.Equal(t, orgs, resultOrgs)
}
