package custom_type

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 自定义类型

// 时间
type Time struct {
	time.Time
}

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

func (t Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	if t.Time.IsZero() {
		output = fmt.Sprint("\"\"")
	}
	return []byte(output), nil
}

func (t Time) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Format(TimeFormat)
}

func ParseTime(timeStr string) (Time, error) {
	dt, err := time.ParseInLocation(TimeFormat, timeStr, time.Local)
	return Time{dt}, err
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = Time{time.Time{}}
		return
	}

	// 指定解析的格式
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Time{now}
	return
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 日期
type Date struct {
	Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", d.Format(DateFormat))
	if d.Time.IsZero() {
		output = fmt.Sprint("\"\"")
	}
	return []byte(output), nil
}

func (d Date) String() string {
	if d.Time.IsZero() {
		return ""
	}
	return d.Format(DateFormat)
}

func ParseDate(timeStr string) (Date, error) {
	dt, err := time.ParseInLocation(DateFormat, timeStr, time.Local)
	return Date{Time{dt}}, err
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*d = Date{Time{time.Time{}}}
		return
	}

	// 指定解析的格式
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(data), time.Local)
	*d = Date{Time{now}}
	return
}
