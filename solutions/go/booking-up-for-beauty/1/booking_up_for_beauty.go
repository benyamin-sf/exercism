package booking

import (
    "time"
    "log"
    "fmt"
)

const DEFAULT_LAYOUT string = "1/2/2006 15:04:05"

func parseDate(layout, dateStr string) (time.Time, error) {
    t, err := time.Parse(layout, dateStr)
    if err != nil {
        return t, fmt.Errorf("Couldn't parse the date: '%w'", err)
    }
    return t, nil
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    t, err := parseDate(DEFAULT_LAYOUT, date)
    if err != nil {
        log.Fatalln(err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	now := time.Now()

    t, err := parseDate(layout, date)
    if err != nil {
        log.Fatalln(err)
    }
    return now.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    var isInAfternoon bool
    layout := "Monday, January 2, 2006 15:04:05"
	t, err := parseDate(layout, date)
    if err != nil {
        log.Fatalln(err)
    }

    appointmentHour := t.Hour()
    if appointmentHour >= 12 && appointmentHour < 18 {
        isInAfternoon = true
    }
    return isInAfternoon
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := parseDate(DEFAULT_LAYOUT, date)
    if err != nil {
        log.Fatalln(err)
    }

    return fmt.Sprintf("You have an appointment on %v.", t.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
    return time.Date(now.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
