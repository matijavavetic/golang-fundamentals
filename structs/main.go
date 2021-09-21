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

	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")

	// Shortcut for the code above
	// Go will automatically resolve a pointer
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}