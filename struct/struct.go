package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//Varsayılan ve boş yapıcı metot
func NewHuman() *Human {
	h := new(Human)
	return h
}

func NewHumanWithParams(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {

	//Nesne örneği oluşturma

	//v1
	/*
		x := Human{
			FirstName: "Anıl",
			LastName:  "Aydın",
			Age:       22,
		}

		fmt.Println(x.FirstName, x.LastName, x.Age)
	*/

	//v2
	/*
		x := new(Human)
		x.FirstName = "Anıl"
		fmt.Println(x.FirstName)
	*/

	//Yapıcı metot kullanımı
	/*
		x := NewHuman()
		x.Age = 22
		fmt.Println(x.Age)
	*/

	//Parametreli yapıcı metot kullanımı

	x := NewHumanWithParams("Anıl", "Aydın", 22)
	//fmt.Println(x.FirstName, x.LastName, x.Age)

	//Veriyi okuyalım
	var data = x.FirstName + " " + x.LastName + " / " + strconv.Itoa(x.Age)
	fmt.Println(data)
}
