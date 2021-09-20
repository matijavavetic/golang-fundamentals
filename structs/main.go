package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	/*
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)	

	// Another way of declaring a struct
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// %+v prints "fieldname: value"
	fmt.Printf("%+v", alex)
	*/

	// Embedding structs
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@mail.com",
			zipCode: 10501,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}