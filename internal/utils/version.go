package utils

var version = "source"

// SetVersion sets the version number for use later in the version command.
func SetVersion(v string) {
	version = v
}

// GetVersion retrieves the version number for use later in the version command.
func GetVersion() string {
	return version
}
