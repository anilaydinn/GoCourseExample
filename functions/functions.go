package main

import "fmt"

//Fonksiyonlar
func main() {
	fmt.Println(add(3, 4, ""))
}

func sayHello(message *string) {
	fmt.Println(*message)
	*message = "Hi Go!"
}

func add(x, y int, z string) int {
	return x + y
}
