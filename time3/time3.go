package main

import (
	"fmt"
	"time"
)

//Tarih ve Zaman operasyonları

func main() {
	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Çalışıyor: %s\n", t)

	fmt.Println("*************************")

	now := time.Now()
	fmt.Printf("Mevcut Zaman: %s", now)

	fmt.Println("*************************")

	fmt.Println("Ay:", now.Month())
	fmt.Println("Gün:", now.Day())
	fmt.Println("Haftanın Günü:", now.Weekday())

	fmt.Println("*************************")

	tomorrow := time.Now().AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	fmt.Println("*************************")

	longFormat := "Monday, January 2, 2018"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "1/2/2020"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}
