package main

import "fmt"

func main() {

	//Konsol : Veri Çıkışı

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)

	fmt.Println("String Length:", stringLength)

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float %.2f\n", float64(aNumber))

	//Golang Placeholders

	fmt.Printf("Data types: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue)

	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T and %T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
