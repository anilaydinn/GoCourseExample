package main

import (
	"log"
	"os"
)

/*
	Dosya silmek (Delete a File)
*/

func main() {
	err := os.Remove("demo.txt")

	if err != nil {
		log.Fatal(err)
	}
}
