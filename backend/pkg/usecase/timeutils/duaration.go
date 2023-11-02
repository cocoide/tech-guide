package timeutils

import "time"

func OneWeek() time.Duration {
	return time.Hour * 24 * 7
}

func OneMonth() time.Duration {
	return time.Hour * 24 * 7 * 30
}

func OneDay() time.Duration {
	return time.Hour * 24
}
