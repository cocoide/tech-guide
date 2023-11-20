package utils

import "time"

func ParseIntSeconds(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}

func NowInJST() time.Time {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
