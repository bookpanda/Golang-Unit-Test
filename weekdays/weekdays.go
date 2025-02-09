package weekdays

import "time"

func GetWeekdaysByYear(year int) ([]time.Time, error) {
	var weekdays []time.Time
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	for d := startDate; d.Year() == year; d = d.AddDate(0, 0, 1) {
		if d.Weekday() >= time.Monday && d.Weekday() <= time.Friday {
			weekdays = append(weekdays, d)
		}
	}
	return weekdays, nil
}