package main

import (
	"fmt"
	"strconv"
)

func main() {

	//kullanici veri olusturma v1

	fmt.Println("kullanici olusturma v1")

	user1 := &user{

		ID:        1,
		firstname: "ahmet can ",
		lastname:  "ceylan",
		username:  "ahmetcanceylan",
		age:       23,
		pay: &payment{
			salary: 3555,
			bonus:  850,
		},
	}

	//fmt.Println(user1)
	fmt.Println(user1.getfullname())
	//fmt.Println(user1.getpayment())
	fmt.Println(user1.pay.salary)
	fmt.Println(user1.pay.bonus)
	fmt.Println("maas 3555 + 850 = "+strconv.FormatFloat(user1.getpayment(), 'g', -1, 64), "TL")
}

// kullanici yapisi
type user struct {
	ID        int
	firstname string
	lastname  string
	username  string
	age       int
	pay       *payment
}

// kullanicinin yapici metodu
func newuser() *user {

	u := new(user)
	u.pay = newpayment()
	return u
}

// odeme yapisi
type payment struct {
	salary float64
	bonus  float64
}

// odemenin yapici metodu
func newpayment() *payment {

	p := new(payment)
	return p
}

//metotlar

func (u user) getfullname() string {

	return u.firstname + "" + u.lastname
}

func (u *user) getusername() string {
	return u.username
}

func (u *user) getpayment() float64 {
	pay := u.pay.salary + u.pay.bonus
	return pay
}
