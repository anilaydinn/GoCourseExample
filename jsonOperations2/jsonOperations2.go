package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSON İşlemleri

//Type Nesneleri

//Name struct'ını tanımlıyoruz.
type Name struct {
	Family   string
	Personal string
}

//Email struct
type Email struct {
	ID      int
	Kind    string
	Address string
}

//Interest struct
type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("***********************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error :", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)

	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func main() {

	person := Person{
		ID:        9,
		FirstName: "Anıl",
		LastName:  "Aydın",
		UserName:  "thracian22",
		Gender:    "true",
		Name:      Name{Family: "Kartelam", Personal: "Anıl"},
		Email: []Email{
			Email{ID: 1, Kind: "work", Address: "anil_aydinn@outlook.com"},
			Email{ID: 2, Kind: "home", Address: "anil_aydin98@hotmail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Go"},
			Interest{ID: 2, Name: "Javascript"},
			Interest{ID: 3, Name: "Java"},
		},
	}

	WriteMessage("Reading Operation Started")

	WriteMessage("Personal Fullname")
	WriteStarLine()

	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personal Email With Index")
	WriteStarLine()

	resEmail := GetPersonEmailAddress(&person, 1)
	WriteMessage(resEmail)
	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personal Email Object with Index")
	WriteStarLine()
	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarLine()

	WriteMessage("Reading operation end.")
	WriteMessage("\n")

	WriteMessage("Writing Operation Started")
	SaveJSON("person.json", person)

	WriteMessage("Writing Operation Ended.")
}
