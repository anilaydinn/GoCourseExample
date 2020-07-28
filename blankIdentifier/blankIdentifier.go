package main

import "fmt"

func main() {

	//Boş Tanımlayıcı ( _ )

	// _, veri := metot()
	// fmt.Println(veri)

	//var _ string = "adada"
	//fmt.Println(_) Çalışmaz

	var _, x, _ int = 0, 9, 0
	fmt.Println(x)

}
