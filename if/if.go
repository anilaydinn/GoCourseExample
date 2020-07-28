package main

import "fmt"

func main() {

	//If

	a := 10
	b := 10

	if b > a {
		fmt.Println("Büyüktür")
	} else if b == a {
		fmt.Println("Eşittir")
	} else {
		fmt.Println("Küçüktür")
	}
}
