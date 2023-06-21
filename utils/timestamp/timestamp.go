// Package timestamp is a package for making it easier to send timestamps in Discord messages
package timestamp

import (
	"fmt"
	"time"
)

// FormatTime returns the time with the given format
func FormatTime(time time.Time, format FormatType) string {
	if format == "" {
		format = FormatTypeShortDateTime
	}
	return fmt.Sprintf("<t:%d:%s>", time.Unix(), format)
}

// FormatType represents the format type of how it will look like in Discord
type FormatType string

const (
	// FormatTypeShortTime will look like the following: 16:20
	FormatTypeShortTime = "t"
	// FormatTypeLongTime will look like the following: 16:20:30
	FormatTypeLongTime = "T"
	// FormatTypeShortDate will look like the following: 20/04/2021
	FormatTypeShortDate = "d"
	// FormatTypeLongDate will look like the following: 20 April 2021
	FormatTypeLongDate = "D"
	// FormatTypeShortDateTime will look like the following: 20 April 2021 16:20 (Default)
	FormatTypeShortDateTime = "f"
	// FormatTypeLongDateTime will look like the following: Tuesday, 20 April 2021 16:20
	FormatTypeLongDateTime = "F"
	// FormatTypeRelative will look like the following: 2 months ago
	FormatTypeRelative = "R"
)
