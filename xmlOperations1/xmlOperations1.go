package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	Sites       []site   `xml:"site"`
	Description string   `xml:",innerxml"`
}

type site struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
	Category    string   `xml:"Category"`
}

func main() {
	file, err := os.Open("sites.xml")

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := ObjectSites{}
	err = xml.Unmarshal(data, &v)

	fmt.Println(v)
}
