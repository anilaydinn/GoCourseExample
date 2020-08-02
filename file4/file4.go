package main

import (
	"log"
	"os"
)

/*
	Dosyanın var olup olmadığını kontrol etmek (Check if File Exists)
*/

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileInfo, err := os.Stat("demo.txt")

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists")
		}
	}
	log.Println("File does exist. File Information: ")
	log.Println(fileInfo)
}
