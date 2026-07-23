package helper

import (
	"time"
	"fmt"
)

func FormatDuration(seconds float64) string {
	duration := time.Duration(seconds) * time.Second

	minutes := int(duration.Minutes())
	remainingSeconds := int(duration.Seconds()) % 60

	return fmt.Sprintf(
		"%02d:%02d",
		minutes,
		remainingSeconds,
	)
}