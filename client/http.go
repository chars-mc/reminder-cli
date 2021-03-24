package client

import (
	"net/http"
	"time"
)

// HTTPClient represents the http client, contains the URI and *http.Client
type HTTPClient struct {
	client     *http.Client
	BackendURI string
}

// NewHTTPClient returns a new HTTPClient
func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		client:     &http.Client{},
	}
}

// Create add a new reminder
func (c HTTPClient) Create(title, message string, duration time.Duration) ([]byte, error) {
	res := []byte("response for create reminder")
	return res, nil
}

// Edit edits a reminder
func (c HTTPClient) Edit(id, title, message string, duration time.Duration) ([]byte, error) {
	res := []byte("response for edit reminder")
	return res, nil
}

// Fetch returns one or more reminders
func (c HTTPClient) Fetch(ids []string) ([]byte, error) {
	res := []byte("response for fetching reminder(s)")
	return res, nil
}

// Delete deletes a reminder
func (c HTTPClient) Delete(ids []string) error {
	return nil
}

// Healthy check if the notifier server is running
func (c HTTPClient) Healthy(host string) bool {
	return true
}
