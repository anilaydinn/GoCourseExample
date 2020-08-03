package main

import (
	"io/ioutil"
	"log"
)

/*
	Hızlı Dosya Yazma (Quick Write to File)
*/

func main() {
	err := ioutil.WriteFile("demo.txt", []byte("Merhaba!"), 0666)

	if err != nil {
		log.Fatal(err)
	}
}
