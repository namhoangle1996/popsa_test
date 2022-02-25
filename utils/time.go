package utils

import "time"

func IsTimeInTheWeekend(i time.Time) bool {
	 return i.Weekday() == time.Saturday || i.Weekday() == time.Sunday
}


