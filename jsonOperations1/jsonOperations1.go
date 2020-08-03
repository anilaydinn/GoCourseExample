package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
		{
			"data":{
				"object":"card",
				"id":"card_4735626194753",
				"first_name":"Anıl",
				"last_name":"Aydın",
				"balance":"54.950"
			}
		}
	`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Println("-------------------")

	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
