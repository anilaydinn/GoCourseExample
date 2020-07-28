package main

import "fmt"

/*
	for döngüsünün range ("aralık") formu ile bir dilim ("slice") veya eşlem ("map") üzerinde dolaşılır.
	Bir dilim üzerinde dolaşılırken her seferinde iki değer dönülür. İlki indis, ikincisi on indisli elemanın bir kopyası.
	Değer dizisinin indisleri veya o indise karşılık gelen değerleri _ ataması ile atlayabilirsiniz.
	Yalnızca indisi kullanmak istiyorsanız ", value" kısmını tamamen çıkarınız.
*/

func main() {

	//Array
	/*
		a := [...]string{"a", "b", "c", "d"}

		for i := range a {
			fmt.Println("Array item", i, "is", a[i])
		}
	*/

	//Map
	baskentler := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	for key, val := range baskentler {
		fmt.Println("Map Item: Capital of", key, "is", val)
	}
}
