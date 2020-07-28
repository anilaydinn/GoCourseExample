package main

import "fmt"

func main() {

	//Convert

	/*
		var myString = "1"
		var x = 10
		var f float32 = 2.0
	*/

	//fmt.Println(myString, x, f)

	//string'den int'e dönüştürme
	//number, _ := strconv.Atoi(myString)

	//fmt.Println("Sonuç: " + strconv.Itoa(number))

	//result := number + 2
	//fmt.Println(result)

	//Casting

	var i int = 55
	//var f1 float64 = i
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
}
