package ranges

import "time"

type Time struct {
	End   time.Time
	Start time.Time
}

func (tr Time) Days() []time.Time {
	days := []time.Time{}
	f := tr.Start
	for f.Before(tr.End) {
		// set to date = from + 1 day
		t := f.AddDate(0, 0, 1)
		days = append(days, t)

		// increment from date with 1
		f = t
	}
	return days
}
