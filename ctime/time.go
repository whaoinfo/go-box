package ctime

import "time"

func CurrentTimestamp() int64 {
	return time.Now().Unix()
}
