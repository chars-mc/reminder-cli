package client

import "net/http"

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
