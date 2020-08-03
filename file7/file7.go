package main

import (
	"io"
	"log"
	"os"
)

/*
	Dosya Kopyalama (Copy a file)
*/

func main() {
	originalFile, err := os.Open("demo.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	//Yeni bir dosya oluştur
	newFile, err := os.Create("./folder/demo_copy.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//Byte'ları kaynaktan hedefe kopyala
	bytesWritten, err := io.Copy(newFile, originalFile)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	//Dosya içeriğini işle
	//Belleği diske boşalt.
	err = newFile.Sync()

	if err != nil {
		log.Fatal(err)
	}
}
