package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Kullanıcı veri oluşturma v1
	//fmt.Println("Kullanıcı oluşturma v1")
	/*
		user1 := &User{
			ID:        1,
			FirstName: "Anıl",
			LastName:  "Aydın",
			UserName:  "thracian22",
			Age:       22,
			Pay: &Payment{
				Salary: 3555,
				Bonus:  850,
			},
		}
	*/
	/*
		fmt.Println(user1)
		fmt.Println(user1.GetFullName())
		fmt.Println(user1.GetPayment())
		fmt.Println("Maaş:" + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64))
	*/

	//Kullanıcı veri oluşturma v2
	fmt.Println("Kullanıcı oluşturma v2")

	user2 := new(User)
	user2.FirstName = "Anıl"
	user2.LastName = "Aydın"
	user2.Age = 22
	user2.UserName = "thracian22"

	//Payment
	user2.Pay = &Payment{
		Salary: 775,
		Bonus:  955,
	}
	fmt.Println(user2)
	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetPayment())
	fmt.Println("Maaş:" + strconv.FormatFloat(user2.GetPayment(), 'g', -1, 64))
}

//Kullanıcı yapısı
type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

//Kullanıcının yapıcı metodu
func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

//Ödeme yapısı
type Payment struct {
	Salary float64
	Bonus  float64
}

//Ödemenin yapıcı metodu
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

//Metotlar
func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
