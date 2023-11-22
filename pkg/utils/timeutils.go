package utils

import (
	"strconv"
	"time"
)

func TimestampInSecondsToISO8601(timestamp string) string {
	blockTimestamp, _ := strconv.ParseInt(timestamp, 10, 64)
	return time.Unix(blockTimestamp, 0).UTC().Format(time.RFC3339)
}
