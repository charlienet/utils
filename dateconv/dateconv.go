package dateconv

import (
	"time"
)

// 日期转换为整数（如:20211222）
func DateToInt(date time.Time) int {
	return date.Year()*10000 + int(date.Month())*100 + date.Day()
}

// 日期转换为字符串
func DateToString(date *time.Time) string { return formatTime(date, "2006-01-02") }

// 时间转换为字符串
func TimeToString(date *time.Time) string { return formatTime(date, "2006-01-02 15:04:05") }

// 日期转换为中文
func DateToChinese(t *time.Time) string { return formatTime(t, "2006年01月02日") }

// 时间转换为中文
func TimeToChinese(t *time.Time) string { return formatTime(t, "2006年01月02日 15:04:05") }

func formatTime(t *time.Time, f string) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(f)
}
