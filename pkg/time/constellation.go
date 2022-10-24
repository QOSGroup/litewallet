package time

import "strings"

func SearchConstellation(str string) (constellation string, idx, minM, minD, maxM, maxD int) {
	if strings.Contains(str, "水瓶") {
		return "水瓶座", 1, 1, 20, 2, 18
	}
	if strings.Contains(str, "双鱼") {
		return "双鱼座", 2, 2, 19, 3, 20
	}
	if strings.Contains(str, "白羊") {
		return "白羊座", 3, 3, 21, 4, 19
	}
	if strings.Contains(str, "金牛") {
		return "金牛座", 4, 4, 20, 5, 20
	}
	if strings.Contains(str, "双子") {
		return "双子座", 5, 5, 21, 6, 21
	}
	if strings.Contains(str, "巨蟹") {
		return "巨蟹座", 6, 6, 22, 7, 22
	}
	if strings.Contains(str, "狮子") {
		return "狮子座", 7, 7, 23, 8, 22
	}
	if strings.Contains(str, "处女") {
		return "处女座", 8, 8, 23, 9, 22
	}
	if strings.Contains(str, "天秤") {
		return "天秤座", 9, 9, 23, 10, 22
	}
	if strings.Contains(str, "天蝎") {
		return "天蝎座", 10, 10, 23, 11, 21
	}
	if strings.Contains(str, "射手") {
		return "射手座", 11, 11, 22, 12, 21
	}
	if strings.Contains(str, "摩羯") {
		return "摩羯座", 12, 12, 22, 1, 19
	}

	return "", 0, 0, 0, 0, 0
}
