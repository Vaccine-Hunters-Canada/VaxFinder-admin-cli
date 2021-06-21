package api

// HTTPOperation abstracts away logic for HTTP Operations for those that implements it
type HTTPOperation interface {
	SetAuthKey(key string)
	SetRequestURLArguments(args []string) error
	GetDetails() (string, string, string)
	GetVerboseResponseFieldNames() []string
	GetResponseAsArray() ([][]string, error)
}
