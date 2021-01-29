package util

import (
	"math/rand"
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

// 导出随机字符串
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 获取星期
func GetWeek() string {
	var datStr string
	day := time.Now().Weekday()
	switch day {
	case 0:
		datStr = "星期天"
	case 1:
		datStr = "星期一"
	case 2:
		datStr = "星期二"
	case 3:
		datStr = "星期三"
	case 4:
		datStr = "星期四"
	case 5:
		datStr = "星期五"
	case 6:
		datStr = "星期六"
	}
	return datStr
}
