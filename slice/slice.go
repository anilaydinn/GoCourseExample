package main

import "fmt"

// Slice

func main() {

	/*
		myArray1 := [...]int{45, 23, 43}
		mySlice1 := myArray1[:]
		fmt.Println(mySlice1)
		mySlice1[0] = 11
		fmt.Println(mySlice1)
		fmt.Println(myArray1)
	*/

	/*
		var colors = []string{"Red", "Green", "Blue"}
		fmt.Println(colors)
		colors = append(colors, "Purple")
		fmt.Println(colors)
		colors = append(colors[1:len(colors)])
		fmt.Println(colors)
		colors = append(colors[:len(colors)-1])
		fmt.Println(colors)
	*/

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 43
	numbers[2] = 87
	numbers[3] = 456
	numbers[4] = 654
	fmt.Println(numbers)

	numbers = append(numbers, 342)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
}
