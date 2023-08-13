package datetime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/*
   @File: gmt.go
   @Author: khaosles
   @Time: 2023/8/13 09:53
   @Desc:
*/

type GMT time.Time

var timeLayout = "2006-01-02 15:04:05"

func (gmt *GMT) UnmarshalJSON(data []byte) error {
	// 加载上海时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(`"`+timeLayout+`"`, string(data), location)
	if err != nil {
		return err
	}
	*gmt = GMT(t)
	return nil
}

func (gmt GMT) MarshalJSON() ([]byte, error) {
	t := time.Time(gmt)
	formatted := fmt.Sprintf(`"%s"`, t.Format(timeLayout))
	return []byte(formatted), nil
}

func (gmt GMT) Value() (driver.Value, error) {
	return time.Time(gmt), nil
}

func (gmt GMT) Scan(value interface{}) error {
	if value == nil {
		gmt = GMT(time.Time{})
		return nil
	}
	if t, ok := value.(time.Time); ok {
		gmt = GMT(t)
		return nil
	}
	return fmt.Errorf("failed to scan CustomTime value")
}

func (gmt GMT) Add(duration time.Duration) GMT {
	return GMT(time.Time(gmt).Add(duration))
}

func (gmt GMT) Sub(gmt2 GMT) time.Duration {
	return time.Time(gmt).Sub(time.Time(gmt2))
}

func (gmt GMT) Format(layout string) string {
	return time.Time(gmt).Format(layout)
}
