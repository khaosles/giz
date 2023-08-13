package datetime

import (
	"fmt"
	"strings"
	"time"
)

/*
   @File: datatime.go
   @Author: khaosles
   @Time: 2023/8/13 09:36
   @Desc:
*/

// Package datetime implements some functions to format date and time.
// Note:
// 1. `format` param in FormatTimeToStr function should be as flow (case no sensitive):
// "yyyy-mm-dd hh:mm:ss"
// "yyyy-mm-dd hh:mm"
// "yyyy-mm-dd hh"
// "yyyy-mm-dd"
// "yyyy-mm"
// "mm-dd"
// "dd-mm-yy hh:mm:ss"
// "yyyy/mm/dd hh:mm:ss"
// "yyyy/mm/dd hh:mm"
// "yyyy/mm/dd hh"
// "yyyy/mm/dd"
// "yyyy/mm"
// "mm/dd"
// "dd/mm/yy hh:mm:ss"
// "yyyymmdd"
// "mmddyy"
// "yyyy"
// "yy"
// "mm"
// "hh:mm:ss"
// "hh:mm"
// "mm:ss"

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyymmdd":            "20060102",
		"mmddyy":              "010206",
		"yyyy":                "2006",
		"yy":                  "06",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"hh:mm":               "15:04",
		"mm:ss":               "04:05",
	}
}

// AddMinute add or sub minute to the time.
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour add or sub hour to the time.
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay add or sub day to the time.
func AddDay(t time.Time, day int64) time.Time {
	return t.Add(24 * time.Hour * time.Duration(day))
}

// AddYear add or sub year to the time.
func AddYear(t time.Time, year int64) time.Time {
	return t.Add(365 * 24 * time.Hour * time.Duration(year))
}

// GetNowDate return format yyyy-mm-dd of current date.
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTime return format hh-mm-ss of current time.
func GetNowTime() string {
	return time.Now().Format("15:04:05")
}

// GetNowDateTime return format yyyy-mm-dd hh-mm-ss of current datetime.
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetTodayStartTime return the start time of today, format: yyyy-mm-dd 00:00:00.
func GetTodayStartTime() string {
	return time.Now().Format("2006-01-02") + " 00:00:00"
}

// GetTodayEndTime return the end time of today, format: yyyy-mm-dd 23:59:59.
func GetTodayEndTime() string {
	return time.Now().Format("2006-01-02") + " 23:59:59"
}

// GetZeroHourTimestamp return timestamp of zero hour (timestamp of 00:00).
func GetZeroHourTimestamp() int64 {
	ts := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", ts)
	return t.UTC().Unix() - 8*3600
}

// GetNightTimestamp return timestamp of zero hour (timestamp of 23:59).
func GetNightTimestamp() int64 {
	return GetZeroHourTimestamp() + 86400 - 1
}

// Strftime convert time to string.
func Strftime(t time.Time, format string, timezone ...string) string {
	tf, ok := timeFormat[strings.ToLower(format)]
	if !ok {
		return ""
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return ""
		}
		return t.In(loc).Format(tf)
	}
	return t.Format(tf)
}

// Strptime convert string to time.
func Strptime(str, format string, timezone ...string) (time.Time, error) {
	tf, ok := timeFormat[strings.ToLower(format)]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not support", format)
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return time.Time{}, err
		}

		return time.ParseInLocation(tf, str, loc)
	}

	return time.Parse(tf, str)
}

// BeginOfMinute return beginning minute time of day.
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute return end minute time of day.
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour return beginning hour time of day.
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour return end hour time of day.
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay return beginning hour time of day.
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay return end time of day.
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek return beginning week, default week begin from Sunday.
func BeginOfWeek(t time.Time, beginFrom ...time.Weekday) time.Time {
	var beginFromWeekday = time.Sunday
	if len(beginFrom) > 0 {
		beginFromWeekday = beginFrom[0]
	}
	y, m, d := t.AddDate(0, 0, int(beginFromWeekday-t.Weekday())).Date()
	beginOfWeek := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	if beginOfWeek.After(t) {
		return beginOfWeek.AddDate(0, 0, -7)
	}
	return beginOfWeek
}

// EndOfWeek return end week time, default week end with Saturday.
func EndOfWeek(t time.Time, endWith ...time.Weekday) time.Time {
	var endWithWeekday = time.Saturday
	if len(endWith) > 0 {
		endWithWeekday = endWith[0]
	}
	y, m, d := t.AddDate(0, 0, int(endWithWeekday-t.Weekday())).Date()
	var endWithWeek = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	if endWithWeek.Before(t) {
		endWithWeek = endWithWeek.AddDate(0, 0, 7)
	}
	return endWithWeek
}

// BeginOfMonth return beginning of month.
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth return end of month.
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear return the date time at the begin of year.
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear return the date time at the end of year.
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// IsLeapYear check if param year is leap year or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// BetweenSeconds returns the number of seconds between two times.
func BetweenSeconds(t1 time.Time, t2 time.Time) int64 {
	index := t2.Unix() - t1.Unix()
	return index
}

// DayOfYear returns which day of the year the parameter date `t` is.
func DayOfYear(t time.Time) int {
	y, m, d := t.Date()
	firstDay := time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
	nowDate := time.Date(y, m, d, 0, 0, 0, 0, t.Location())

	return int(nowDate.Sub(firstDay).Hours() / 24)
}

// IsWeekend checks if passed time is weekend or not.
// Deprecated Use '== Weekday' instead
func IsWeekend(t time.Time) bool {
	return time.Saturday == t.Weekday() || time.Sunday == t.Weekday()
}

// NowTime return current datetime.
func NowTime(timezone ...string) time.Time {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return time.Time{}
		}
		return time.Now().In(loc)
	}
	return time.Now()
}

// NowDateOrTime return current datetime with specific format and timezone.
func NowDateOrTime(format string, timezone ...string) string {
	tf, ok := timeFormat[strings.ToLower(format)]
	if !ok {
		return ""
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return ""
		}

		return time.Now().In(loc).Format(tf)
	}

	return time.Now().Format(tf)
}

// Timestamp return current second timestamp.
func Timestamp(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		t = t.In(loc)
	}

	return t.Unix()
}

// TimestampMilli return current mill second timestamp.
func TimestampMilli(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}
		t = t.In(loc)
	}

	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}

// TimestampMicro return current micro second timestamp.
func TimestampMicro(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}
		t = t.In(loc)
	}

	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Microsecond)
}

// TimestampNano return current nano second timestamp.
func TimestampNano(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}
		t = t.In(loc)
	}

	return t.UnixNano()
}

// ToTimezone return target timezone datetime if timezone exist
func ToTimezone(t time.Time, timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return t
	}
	return t.In(loc)
}
