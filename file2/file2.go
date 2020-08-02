package main

import (
	"fmt"
	"log"
	"os"
)

/*
	Dosya Bilgisi Almak (Get File Info)
*/

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {

	//Dosya bilgisini döndürür eğer dosya yoksa hata döner.
	fileInfo, err := os.Stat("demo.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name :", fileInfo.Name())
	fmt.Println("Size in bytes :", fileInfo.Size())
	fmt.Println("Permissions :", fileInfo.Mode())
	fmt.Println("Last Modified :", fileInfo.ModTime())
	fmt.Println("Is Directory :", fileInfo.IsDir())
	fmt.Println("System Interface Type :", fileInfo.Sys())
	fmt.Printf("System Info : %+v\n\n", fileInfo.Sys())
}
