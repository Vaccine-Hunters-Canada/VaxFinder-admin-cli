package utils

import (
	"strconv"
)

// CoalesceString returns empty string on null, otherwise returns the value
func CoalesceString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

// CoalesceInt returns empty string on null, otherwise returns the value as string
func CoalesceInt(value *int) string {
	if value == nil {
		return ""
	}
	return strconv.Itoa(*value)
}
