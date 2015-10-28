package am

import "time"

// TimeWithLocation returns a time.Time using the given location,
// while using the time values from the given time (day, hour, minutes, etc.)
func TimeWithLocation(l *time.Location, t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), l)
}

// TimeWithLocations returns a slice of time.Time using the given location,
// while using the time values from the given time (day, hour, minutes, etc.)
func TimesWithLocation(l *time.Location, t ...time.Time) []time.Time {
	s := []time.Time{}
	for i := range t {
		s = append(s, TimeWithLocation(l, t[i]))
	}
	return s
}
