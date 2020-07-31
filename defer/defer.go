package main

import "fmt"

//Fonksiyonlar : defer

var isConnected bool = false

func main() {
	fmt.Printf("Connection open: %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection open: %v\n", isConnected)
}

func databaseProcessing() {
	connect()
	fmt.Println("Defering disconnect!")
	defer disconnect()
	fmt.Printf("Connection open: %v\n", isConnected)
	fmt.Println("Doing something")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to database!")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")
}
