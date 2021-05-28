package api

// HTTPOperation abstracts away logic for HTTP Operations for those that implements it
type HTTPOperation interface {
	GetDetails() (string, string, string)
	SetRequestURLArguments(args []string) error
	GetVerboseResponseFieldNames() []string
	GetResponseAsArray() ([][]string, error)
}
