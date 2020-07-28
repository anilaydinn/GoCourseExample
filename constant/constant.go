package main

import "fmt"

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE    Brand = "Google"
	DIJIBIL   Brand = "Dijibil"
)

func PrintBrand(brand Brand) {
	fmt.Println(brand)
}

func main() {

	//Constant
	PrintBrand(FACEBOOK)
	PrintBrand(DIJIBIL)
}
