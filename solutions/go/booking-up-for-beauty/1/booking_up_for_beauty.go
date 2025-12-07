package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layoutFrom := "1/2/2006 15:04:05"
	t, err := time.Parse(layoutFrom, date); if err != nil {return time.Time{}}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date); if err != nil { return false }
	return time.Now().After(t)
}


// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date); if err != nil { return false }
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %02d:%02d.",
		t.Weekday(), t.Month(), t.Day(), t.Year(),
		t.Hour(), t.Minute(),
	)
}


// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
