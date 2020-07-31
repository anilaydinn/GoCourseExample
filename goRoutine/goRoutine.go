package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//basit gorouting kullanımı
	//go xFunc()
	//time.Sleep(100 * time.Millisecond)

	runtime.GOMAXPROCS(1)
	go xFunc()
	time.Sleep(100 * time.Millisecond)
}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		go fmt.Println(string(l))
	}
}
