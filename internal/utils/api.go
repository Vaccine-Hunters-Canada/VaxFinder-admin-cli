package utils

import (
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"vf-admin/internal/api"
)

var baseURL string

// SetBaseURL sets the base URL for use later in HTTP requests.
func SetBaseURL(b string) {
	baseURL = b
}

// GetBaseURL retrieves the base URL for use later in HTTP requests.
func GetBaseURL() string {
	return baseURL
}

// GetDefaultSecurityProvider returns the default security provider for the VaxFinder API.
func GetDefaultSecurityProvider(key string) (*securityprovider.SecurityProviderApiKey, error) {
	return securityprovider.NewSecurityProviderApiKey("header", "Authorization", "Bearer "+key)
}

// GetAPIClientFromKey returns the API client to be used to make requests with the VaxFinder server.
func GetAPIClientFromKey(key string) (*api.ClientWithResponses, error) {
	provider, err := GetDefaultSecurityProvider(key)
	if err != nil {
		return nil, err
	}

	return api.NewClientWithResponses(
		"https://vax-availability-api.azurewebsites.net",
		api.WithBaseURL(baseURL),
		api.WithRequestEditorFn(provider.Intercept))
}

// GetAPIClient returns the API client to be used to make requests with the VaxFinder server.
func GetAPIClient() (*api.ClientWithResponses, error) {
	return api.NewClientWithResponses(
		"https://vax-availability-api.azurewebsites.net",
		api.WithBaseURL(baseURL))
}
