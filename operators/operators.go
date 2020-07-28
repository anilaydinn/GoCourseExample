package main

import "fmt"

func main() {
	//Operatörler

	a := 10
	b := 20

	total := a + b

	fmt.Println("total: ", total)

	total = total - 5

	fmt.Println("total: ", total)

	total *= 20

	fmt.Println("total: ", total)

	total = 10 / 2

	fmt.Println("total: ", total)

	total = a % b

	fmt.Println("total: ", total)

	total++
	fmt.Println("total: ", total)

	total--
	fmt.Println("total: ", total)
}
