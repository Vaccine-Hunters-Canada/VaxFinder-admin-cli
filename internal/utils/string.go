package utils

import (
	"strconv"
	"time"

	"github.com/nleeper/goment"
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

// "Goment's parser is strict, and defaults to being accurate over forgiving."
// As such, check if goment's parser worked, if error is returned, use original date
func GetFromNow(date time.Time) string {
	g, err := goment.New(date)
	fromNow := g.FromNow()

	if err == nil {
		return string(fromNow)
	} else {
		return date.String()
	}

}
