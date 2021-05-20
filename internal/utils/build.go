package utils

var (
	version string
	tag     string
	date    string
)

// SetBuildInfo sets the build information for use later in the version command.
func SetBuildInfo(v, t, d string) {
	version = v
	tag = t
	date = d
}

// GetBuildInfo retrieves the build information for use later in the version command.
func GetBuildInfo() (string, string, string) {
	return version, tag, date
}
