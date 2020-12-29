package logit

import (
	"time"

	"github.com/vjeantet/jodaTime"
)

// convertDateFormat converts date-time format from joda style into Go style.
func convertDateFormat(format string) string {
	date := time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)
	return jodaTime.Format(format, date)
}
