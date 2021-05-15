package utils

import (
	"github.com/theckman/yacspin"
	"strings"
	"time"
)

// GetDefaultSpinner returns the default spinner
func GetDefaultSpinner(suffix, startMsg, stopMsg, stopFailMsg string) (*yacspin.Spinner, error) {
	cfg := yacspin.Config{
		Frequency:       120 * time.Millisecond,
		HideCursor:      true,
		ColorAll:        true,
		Colors:          []string{"fgCyan"},
		CharSet:         yacspin.CharSets[59],
		Suffix:          suffix,
		SuffixAutoColon: true,
		Message:         startMsg,
		StopMessage:     stopMsg,
		StopColors:      []string{"fgGreen", "bold"},
		StopFailMessage: stopFailMsg,
		StopFailColors:  []string{"fgRed", "bold"},
	}

	return yacspin.New(cfg)
}

// GetDefaultSpinnerForHTTPOp returns the default spinner for a HTTP operation
func GetDefaultSpinnerForHTTPOp(op, oped, entity string) (*yacspin.Spinner, error) {
	suffix := " " + strings.Title(strings.ToLower(op)) + " " + strings.ToLower(entity)
	startMsg := "Making request to server"
	stopMsg := "Successfully " + strings.ToLower(oped) + " " + strings.ToLower(entity)
	stopFailedMsg := "Failed to " + strings.ToLower(op) + " " + strings.ToLower(entity)
	return GetDefaultSpinner(suffix, startMsg, stopMsg, stopFailedMsg)
}
