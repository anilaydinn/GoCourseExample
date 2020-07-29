package main

import "fmt"

func main() {
	/*
		//1.
		myMap := make(map[int]string)
		fmt.Println(myMap)
		myMap[43] = "Foo"
		myMap[12] = "Bar"
		fmt.Println(myMap)
		//key value pair
	*/

	//2.
	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	//şehir listesinde ANT anahtar adına sahip veriyi elde et.
	antalya := states["ANT"]
	fmt.Println("Seçili şehir", antalya)

	//ANK anahtar adına sahip veriyi sil
	delete(states, "ANK")
	fmt.Println(states)

	states["ERZ"] = "Erzurum"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//states map nesnesinin elemanı adedince kapasiteye sahip bir key lisleti oluştur.
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	//key listesindeki index değerlerine göre states nesnesinde bulunan şehirleri yazdır
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
