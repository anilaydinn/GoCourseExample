package main

import "fmt"

func main() {

	//break & continue

	//break
	/*
		i := 0
		for {

			if i >= 3 {
				break
			}

			fmt.Println("i:", i)
			i++
		}
	*/

	//continue
	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println("Çıktı:", i)
	}
	fmt.Println("İşlem devam..")
}
