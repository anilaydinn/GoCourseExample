package main

import (
	"log"
	"os"
)

/*
	Okuma ve Yazma izinlerini kontrol etmek (Check read and write permission)
*/

func main() {

	//Yazma izinlerini test
	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			log.Println("Hata: Yazma izni reddedildi.")
		}
	}
	file.Close()
}
