package date_utils

import "time"

func GetNowString() string {
	now := time.Now().UTC()
	return now.Format("2006-01-02T15:04:05Z")
}