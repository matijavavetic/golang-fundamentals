package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	/*
	Since both englishBot and spanishBot have a function getGreeting() string
	that is declared in the "bot" interface, they are now a member
	of type "bot", which is why we can throw them into this printGreeting function
	which accepts the type "bot" as a parameter
	*/
	printGreeting(sb)
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// If we are not using the variable from a receiver
// we can just specify the type without the variable
func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}