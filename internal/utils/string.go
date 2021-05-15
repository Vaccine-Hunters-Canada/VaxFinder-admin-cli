package utils

// CoalesceString returns empty string on null, otherwise returns the value
func CoalesceString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
