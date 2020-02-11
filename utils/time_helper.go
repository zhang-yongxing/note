package utils

import "time"

func TimestampToDatetime(f int64) time.Time{
	tm := time.Unix(f, 0)
	return tm
}

func DatetimeToTimestamp(time time.Time) int64{
	tm := time.Unix()
	return tm
}

