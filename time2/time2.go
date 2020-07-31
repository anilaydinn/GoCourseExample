package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
	time.Sleep(2 * time.Second)
	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
}

/*
	Unix Time Nedir ?

	Unix Zaman, 1 Ocak 1970 (01/01/1970) den beridir geçen saniye sayısına denilen sayısal veri tipidir.
	[1] 9 Eylül 2001 (09/09/2001) saat 04:46:40 itibari ile Unix Zaman 10 haneye yükselmiştir.
	Ve 20 Kasım 2286 (20/11/2286) saat 19:46:40 da 11 haneye yükselecektir.

*/
