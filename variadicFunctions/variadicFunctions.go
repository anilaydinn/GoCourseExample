package main

import "fmt"

//Fonksiyonlar : Değişken Sayıda Parametre Alan Fonksiyonlar (Variadic Functions)

func main() {
	//sayHello("Merhaba", "Go", "Developers", ".")
	result := add(4, 5, 34, 3, 4, 3)
	fmt.Println(result)
}

func sayHello(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}
