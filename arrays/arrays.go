package main

import "fmt"

//Diziler

func main() {

	//Basit bir dizi
	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54
	fmt.Println(myArray1)

	//Renk dizisi
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])
	colors[1] = "Bordo"
	fmt.Println(colors[1])

	var numbers = [5]int{4, 6, 7, 3, 66}
	fmt.Println(numbers)
	fmt.Println("Number of numbers:", len(numbers))

	myArray2 := [...]int{4, 3, 5, 6, 7, 332}
	myArray2[3] = 555
	fmt.Println(myArray2)

	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "Mercedes"
	cars[2] = "Jaguar"
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	for i, value := range cars {
		fmt.Println("i=", i, "cars is", value)
	}

}
