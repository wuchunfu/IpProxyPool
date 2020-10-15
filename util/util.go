package util

import (
	"regexp"
	"strconv"
	"time"
)

func FormatDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ExtractSpeed(text string) int64 {
	reg := regexp.MustCompile(`\[1-9\]\d\*\\.\?\d\*`)
	temp := reg.FindString(text)
	if temp != "" {
		speed, _ := strconv.ParseInt(temp, 10, 64)
		return speed
	}
	return -1
}
