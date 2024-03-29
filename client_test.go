package centralclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	pc "github.com/t11e/go-pebbleclient"
)

func newTestClient(host string) (Client, error) {
	c, err := pc.NewHTTPClient(pc.Options{
		Host:        host,
		ServiceName: "central",
	})
	if err != nil {
		return nil, err
	}
	return New(c)
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
		w.Header().Set("Content-Type", "application/json")
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

	_, err = client.GetApplicationByKey("frobnitz")
	assert.Error(t, err)

	badKeyErr, ok := err.(*BadAPIKey)
	if !assert.True(t, ok) {
		return
	}
	assert.Equal(t, "frobnitz", badKeyErr.Key)
}

func TestClient_GetOrganization_found(t *testing.T) {
	org := &Organization{
		Id:    1,
		Title: "The Illuminati",
		Realm: "illuminati",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/api/central/v1/organizations/1", req.URL.Path)
		w.Header().Set("Content-Type", "application/json")
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

	_, err = client.GetOrganization(1)
	assert.Error(t, err)
	notFound, ok := err.(*NoSuchOrganization)
	if !assert.True(t, ok) {
		return
	}
	assert.Equal(t, 1, notFound.Id)
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
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(orgs)
	}))

	client, err := newTestClient(hostFromUrl(server.URL))
	assert.NoError(t, err)

	resultOrgs, err := client.GetOrganizations()
	assert.NoError(t, err)
	assert.Equal(t, orgs, resultOrgs)
}
