package api

import "moul.io/http2curl"

// HTTPOperation abstracts away logic for HTTP Operations for those that implements it
type HTTPOperation interface {
	SetAuthKey(key string)
	SetRequestURLArguments(args []string) error
	GetDetails() (string, string, string)
	GetVerboseResponseFieldNames() []string
	GetResponseAsArray() ([][]string, error)
	GetAsCurlCommand(withKey bool) (*http2curl.CurlCommand, error)
}
