package main

import (
	"fmt"
	"time"
)

//Tarih ve Zaman Operasyonları

func main() {

	x := fmt.Println

	xTime := time.Date(1071, 10, 11, 20, 23, 0, 0, time.UTC)
	now := time.Now()
	x(now)

	x("--------------")
	/*
		x(now.Year())
		x(now.Month())
		x(now.Day())
		x(now.Hour())
		x(now.Minute())
		x(now.Second())
		x(now.Nanosecond())
		x(now.Location())
	*/

	// Tarih karşılaştırma ya da kontrol ya da test
	/*
		x(xTime.Before(now))
		x(xTime.After(now))
		x(xTime.Equal(now))
	*/

	diff := now.Sub(xTime)
	x(diff)

	x(diff.Hours())
	x(diff.Minutes())
	x(diff.Seconds())
	x(diff.Nanoseconds())
}
