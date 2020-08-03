package main

import (
	"log"
	"os"
)

/*
	Byte'ları Bir Dosyaya Yazın (Write Bytes to a File)
*/

func main() {
	//Demo.txt dosyasını sadece yazılabilir bir dosya olarak aç.
	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Bu dosyaya yazdık!\n")

	bytesWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
