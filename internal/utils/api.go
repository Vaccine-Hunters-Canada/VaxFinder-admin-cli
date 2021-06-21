package utils

var baseURL string

// SetBaseURL sets the base URL for use later in HTTP requests.
func SetBaseURL(b string) {
	baseURL = b
}

// GetBaseURL retrieves the base URL for use later in HTTP requests.
func GetBaseURL() string {
	return baseURL
}
