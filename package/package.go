package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	log "github.com/koding/logging"
)

/*

	Paket Nedir ?

	Bir araya toplanmış(bundled) fonksiyonlar topluluğudur. Farklı birçok nesneyi de kendi içerisinde barındırabilir.

	- Her Go programı paketlerden oluşur.
	- Programlar main paketinde çalışmaya başlar.

	Faydaları nelerdir?

	- Namespace, yani ilgili fonksiyonları kapsayan ortak isim uzayı sağlar ve proje ya da, sistemdeki diğer fonksiyonlarla karışmamayı sağlar.
	- Kod organizasyonu sağlar.
	- Derleyiciyi hızlandırır. Örneğin fmt her kullanışımızda derleyici tarafından derlenmez.
*/

func main() {

	//fmt
	//fmt.Println("Bu bir örnek")

	//rand
	//fmt.Println("My favorite number is", rand.Intn(10))

	//strings
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("unit_test", "unit"))
	fmt.Println(strings.HasSuffix("dosya.rar", "rar"))
	fmt.Println(strings.Index("test", "e"))

	//color
	color.Red("Error Message!")

	//logging.Info("Uygulama sonu.")

	log.Info("Uygulama sonu.")
}
