package distance

import (
	"log"
	"math"
)

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

// 计算两个经纬度之间的距离  (Haversine公式)
func Haversine(lon1 float64, lat1 float64, lon2 float64, lat2 float64) float64 {

	log.Println("[Haversine]", lon1, " ", lat1, " ", lon2, "", lat2)

	rlon1 := Degrees2Radians(lon1)
	rlat1 := Degrees2Radians(lat1)
	rlon2 := Degrees2Radians(lon2)
	rlat2 := Degrees2Radians(lat2)

	// haversine公式,计算两点之间的距离
	dlon := rlon2 - rlon1
	dlat := rlat2 - rlat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))
	r := float64(6371) // 地球平均半径，单位为公里
	return c * r
}

// 角度转换为弧度
// radians = degrees * (Math.PI/180);
func Degrees2Radians(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// 弧度转换为角度
// degrees = radians * (180/Math.PI);
func Radians2Degrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
