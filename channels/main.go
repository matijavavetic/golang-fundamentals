package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	// This will create a channel where we can only communicate
	// through strings
	c := make(chan string)

	for _, link := range links {
		// With keyword "go" we create an additional Go routine
		// For each call of the checkLink function
		// This code will have a common bug that occurs with main and child go routines
		// After main routine finishes with looping it will have nothing else to do
		// So it will exit the program while the child routines still haven't finished their tasks
		// go checkLink(link)

		// To fix the bug above, we use channels
		go checkLink(link, c)
	}

	// When we wait for message from a channel we create a blocking call
	// Main routine will go to sleep and wait for a message from a channel
	// When the message comes the main routine will printout the message and exit the program
	// Therefore only one request will be executed
	//fmt.Println(<- c)

	/* To fix the behavior mentioned above
	for i := 0; i < len(links); i++ {
		// Remember - BLOCKING CALL
		//fmt.Println(<- c)
	}
	*/

	// With this loop syntax, we are waiting on a value from a channel
	// When value comes it will be assigned to the variable "l"
	for l := range c {
		// Function literal == Lambdas in C#
		go func(link string) {
			// Pausing the current routine
			time.Sleep(5 * time.Second)
			// Repeating routines
			checkLink(link, c)
		} (l) // variable from outer scope to use in function literal
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		//c <- "Might be down I think"
		c <- link
		return
	}

	//c <- "Yep it's up"

	fmt.Println(link, "is up")
	c <- link
}