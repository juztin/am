package am

import "time"

// SetLocation updates the location of the given time,
// while preserving the time values (day, hour, minutes, etc.)
func SetLocation(l *time.Location, t *time.Time) {
	if l == nil {
		l = time.UTC
	}
	_, z1 := t.Zone()
	ll := t.Location()
	*ll = *l
	_, z2 := t.Zone()
	d := int64(time.Second) * int64(z1-z2)
	*t = t.Add(time.Duration(d))
}

// SetLocation updates the location of the given times,
// while preserving the time values (day, hour, minutes, etc.)
func SetLocations(l *time.Location, t ...*time.Time) {
	for i := range t {
		SetLocation(l, t[i])
	}
}
