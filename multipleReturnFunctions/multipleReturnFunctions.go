package main

import "fmt"

//Fonksiyonlar: Çoklu Sonuç Dönmek

func main() {

	//swap
	a, b := swap("Anıl", "Aydın")

	fmt.Println(a, b)

	numTerms, sum := add(3, 4, 5)
	fmt.Println("Added:", numTerms, "terms for a total of sum:", sum)
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}
