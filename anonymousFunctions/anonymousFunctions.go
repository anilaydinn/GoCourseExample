package main

import "fmt"

//Fonksiyonlar : Anonim Fonksiyonlar

func main() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}

	numTerms, sum := addFunc(2, 5)
	fmt.Println("Added:", numTerms, "terms for total of", sum)
}
