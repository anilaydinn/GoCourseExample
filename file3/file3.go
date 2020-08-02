package main

import (
	"log"
	"os"
)

/*
	Yeniden İsimlendirme ve Taşıma (Rename and Move a File)
*/

func main() {

	originalPath := "demo.txt"
	newPath := "./moved/test.txt"
	err := os.Rename(originalPath, newPath)

	if err != nil {
		log.Fatal(err)
	}
}
