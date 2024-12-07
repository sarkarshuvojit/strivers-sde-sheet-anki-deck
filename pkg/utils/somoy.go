package utils

import "time"

func NowInFilesafeFormat() string {
	currentTime := time.Now()

	timestamp := currentTime.Format("2006-01-02_15-04-05")
	return timestamp
}
