package helpers

import "time"

func DateFormat(d string, l string) string  {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	var r string

	if l == "US" {
		t, _ := time.Parse(layoutISO, d)
		r = t.Format(layoutUS)
	} else if l == "ISO" {
		t, _ := time.Parse(layoutISO, d)
		r = t.Format(layoutISO)
	}

	return r
}
