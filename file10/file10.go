package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//geçici klasor oluştur
	tempDirPath, err := ioutil.TempDir("", "geciciKlasor")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Geçici klasör oluşturuldu:", tempDirPath)

	//gecici dosya oluştur.
	tempFile, err := ioutil.TempFile(tempDirPath, "geciciDosya.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Geçici dosya oluşturuldu:", tempFile.Name())

	//close file
	err = tempFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	//Silme
	err = os.Remove(tempFile.Name())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s dosyası silindi.", tempFile.Name())
	}

	err = os.Remove(tempDirPath)

	if err != nil {
		log.Fatal(err)
	}
}
