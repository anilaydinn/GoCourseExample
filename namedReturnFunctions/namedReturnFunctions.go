package main

import "fmt"

//Fonksiyonlar : Adlandırılmış Geri Dönüşler

/*
	Go'nun sonuç değerleri aynı bir değişken gibi adlandırılabilir.
	Böyle yapıldığında değişken sanki fonksiyonun başında tanımlanmış gibi muamele görür.

	Sonuçlar manalarına, işlevlerine göre adlandırılmalıdır.
	Argümansız bir return deyimi adlandırılmış sonuç değerlerini döndürür. Bu "çıplak" return olarak bilinir.
	Çıplak return deyimleri, aynı örnekteki gibi, sadece kısa fonksiyonlarda kullanılmalıdır.
	Uzun fonksiyonlarda kullanılırsa kodun okunabilirliğini azaltırlar.
*/

func main() {

	//1.
	fmt.Println(split(17))

	//2.
	numTerms, sum := add(1, 3)
	println("Added:", numTerms, "terms for a total of", sum)
}

//1.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//2.
func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
