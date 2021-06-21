package api

import (
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"vf-admin/internal/utils"
)

// GetDefaultSecurityProvider returns the default security provider for the VaxFinder API.
func GetDefaultSecurityProvider(key string) (*securityprovider.SecurityProviderApiKey, error) {
	return securityprovider.NewSecurityProviderApiKey("header", "Authorization", "Bearer "+key)
}

// GetAPIClientFromKey returns the API client to be used to make requests with the VaxFinder server.
func GetAPIClientFromKey(key string) (*ClientWithResponses, error) {
	provider, err := GetDefaultSecurityProvider(key)
	if err != nil {
		return nil, err
	}

	if key == "" {
		return GetAPIClient()
	}

	return NewClientWithResponses(
		utils.GetBaseURL(),
		WithBaseURL(utils.GetBaseURL()),
		WithRequestEditorFn(provider.Intercept))
}

// GetAPIClient returns the API client to be used to make requests with the VaxFinder server.
func GetAPIClient() (*ClientWithResponses, error) {
	return NewClientWithResponses(
		utils.GetBaseURL(),
		WithBaseURL(utils.GetBaseURL()))
}
