package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	Username  string   `xml:"username"`
}

type jsonPerson struct {
	ID        int
	FirstName string
	LastName  string
	Username  string
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string {
	return fmt.Sprintf("\t ID: %d - FirstName: %s - LastName: %s - Username: %s\n", p.ID, p.FirstName, p.LastName, p.Username)
}

func main() {
	xmlFile, err := os.Open("Employees.xml")

	if err != nil {
		fmt.Println("Opening file error:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("Read error:", err)
	}

	var c Company
	xml.Unmarshal(xmlData, &c)

	fmt.Println(c.Persons)

	//Convert to JSON
	var person jsonPerson
	var persons []jsonPerson

	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.Username = value.Username

		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./Employees.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

}
