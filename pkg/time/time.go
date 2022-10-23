package time

import (
	"time"

	"yuxiaole/backend-common/util"
)

const DateForm = "2006-01-02"
const TimeForm = "2006-01-02 15:04:05"

func GetTimeFromStrDate(date string) (year, month, day int) {
	d, err := time.Parse(DateForm, date)
	if err != nil {
		return 0, 0, 0
	}
	year = d.Year()
	month = int(d.Month())
	day = d.Day()
	return
}

func GetZodiac(year int) (zodiac string) {
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

func GetAge(year int) (age int) {
	if year <= 0 {
		age = -1
	}
	nowyear := time.Now().Year()
	age = nowyear - year
	return
}

func GetConstellation(month, day int) (star string) {
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
		star = "摩羯座"
	}

	return star
}

func IsSameConstellationClass(c1, c2 string) (bool, int, string) {
	if (c1 == "白羊座" || c1 == "狮子座" || c1 == "射手座") && (c2 == "白羊座" || c2 == "狮子座" || c2 == "射手座") {
		return true, 1, "火象星座"
	}
	if (c1 == "摩羯座" || c1 == "金牛座" || c1 == "处女座") && (c2 == "摩羯座" || c2 == "金牛座" || c2 == "处女座") {
		return true, 2, "土象星座"
	}
	if (c1 == "天秤座" || c1 == "水瓶座" || c1 == "双子座") && (c2 == "天秤座" || c2 == "水瓶座" || c2 == "双子座") {
		return true, 3, "风象星座"
	}
	if (c1 == "巨蟹座" || c1 == "天蝎座" || c1 == "双鱼座") && (c2 == "巨蟹座" || c2 == "天蝎座" || c2 == "双鱼座") {
		return true, 4, "水象星座"
	}
	return false, 0, ""
}

// 是否成年
func IsAdult(bornDate string) (bool, error) {

	if len(bornDate) == 0 { //如果没有传入
		return true, nil
	}
	d, err := time.ParseInLocation(DateForm, bornDate, time.Local)
	if err != nil {
		return false, err
	}

	age := util.CalculateAge(d)
	if age < 18 {
		return false, nil
	}
	return true, nil
}
