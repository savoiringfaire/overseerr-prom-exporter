package overseerr

import "net/http"

// Overseerr provides a nice go interface to the Overseerr API.
type Overseerr struct {
	// The API key to use to authenticate with Overseerr.
	APIKey string
	// The http client to use when making API requests to Overseerr.
	HTTPClient *http.Client
	// API base URL
	BaseURL string
}

// New creates a new instance of Overseerr with sensible defaults.
func New(BaseURL string, APIKey string) Overseerr {
	return Overseerr{
		APIKey:     APIKey,
		HTTPClient: &http.Client{},
		BaseURL:    BaseURL,
	}
}
