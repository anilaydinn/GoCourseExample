package main

import (
	"fmt"
	"strconv"
)

func CarExecute(c Carface) {
	fmt.Println("\n")
	fmt.Println("Araç Bilgi : \n", c.Information())
	fmt.Println("\n")

	msg := ""
	isRun := c.Run()
	if isRun {
		msg = "Çalışıyor"
	} else {
		msg = "Çalışmıyor"
	}
	fmt.Println("Araç " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "durdu"
	} else {
		msg = "duramadı! Fren tutmuyor..."
	}
	fmt.Println("Araç " + msg + ".")

}

func main() {

	//Ferrari
	ferr := NewFerrari("Ferrari")
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.4
	ferr.Special = true

	//fmt.Println(ferr.Information())

	merc := NewMercedes("Mercedes")
	merc.Model = "CLX"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 0.4
	//fmt.Println(merc.Information())

	CarExecute(ferr)
	CarExecute(merc)
}

//Interface

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base structs

type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

//Ferrari

type Ferrari struct {
	Car //composition
	SpecialProduction
}

func NewFerrari(brand string) *Ferrari {
	ferr := new(Ferrari)
	ferr.Brand = brand
	return ferr
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add + "\n"
	}
	return ret
}

//Lamborghini
type Lamborghini struct {
	Car
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (x *Lamborghini) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add
	}
	return ret
}

//Mercedes
type Mercedes struct {
	Car
}

func NewMercedes(brand string) *Mercedes {
	merc := new(Mercedes)
	merc.Brand = brand
	return merc
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color : " + x.Color + "\n" + "\t" + "Speed : " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price : " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	return ret
}
