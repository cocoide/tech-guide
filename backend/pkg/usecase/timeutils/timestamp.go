package timeutils

import (
	"time"
)

func NowInJST() time.Time {
	return time.Now().In(JSTLoc())
}

func TimestampInJST() int64 {
	return time.Now().In(JSTLoc()).Unix()
}

func UnixToJST(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).In(JSTLoc())
}

func JSTLoc() *time.Location {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return loc
}
