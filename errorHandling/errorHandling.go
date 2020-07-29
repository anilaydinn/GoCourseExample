package main

import (
	"fmt"
	"os"
)

//Hata Yönetimi : Error Handling / Exception Handling

func main() {

	_, err := os.Open("abc.rar")

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
}
