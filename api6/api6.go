package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	//String birleştirme işlemi
	var builder bytes.Buffer
	builder.WriteString("Kodlab Yayınevinden çıkardığımız yazılımcılar için ileri seviye T-SQL....\n")
	builder.WriteString("704 Sayfa\n")
	builder.WriteString("ISBN: 9.786.45163.542\n")
	builder.WriteString("Fiyat: 37 TL\n")
	builder.WriteString("Boyut: 15 x 21\n")
	builder.WriteString("2. Baskı\n")

	uri := "www.anilaydin.me"
	page := Page{
		Title:   "Kitap: İleri Seviye T-SQL Programlama",
		Author:  "Cihan Özhan",
		Header:  "İleri Seviye T-SQL Programlama",
		Content: builder.String(),
		URI:     "http://" + uri,
	}

	t, _ := template.ParseFiles("page.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
