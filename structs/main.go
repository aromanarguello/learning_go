package main

import (
	"fmt"
)


type contactInfo struct {
	email string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName: "Roman",
		contactInfo: contactInfo{
			email: "aroman@gmail.com",
		},
	}
	
	alex.updateName("Alejandro")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}