package utils

import (
	"time"
)

func ZeroTime() time.Time {
	return time.Time{}
}

// DayStart the start of today
func DayStart(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// MonthStart start of this month
func MonthStart(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// YearStart start of this year
func YearStart(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// NDaysAgo N天前
func NDaysAgo(t time.Time, n int) time.Time {
	return t.AddDate(0, 0, 0-n)
}

// NDaysAfter N天后
func NDaysAfter(t time.Time, n int) time.Time {
	return t.AddDate(0, 0, n)
}

// NYearsAgo N年前
func NYearsAgo(t time.Time, n int) time.Time {
	return t.AddDate(0-n, 0, 0)
}

// NYearsAfter N年后
func NYearsAfter(t time.Time, n int) time.Time {
	return t.AddDate(n, 0, 0)
}

// DayCountOfYearMonth 计算某一年某一个月的天数
func DayCountOfYearMonth(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Equal(DayStart(now)) || (t.After(DayStart(now)) && t.Before(NDaysAfter(t, 1)))
}

// GetZodiac zodiac
func GetZodiac(t time.Time) (zodiac string) {
	year := t.Year()
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func GetConstellation(t time.Time) (star string) {
	_, month, day := t.Date()
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}

	return star
}
