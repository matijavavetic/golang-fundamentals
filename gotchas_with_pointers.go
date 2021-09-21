package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "you"}

	updateSlice(mySlice)

	/*
	Even though we didn't use pointer in the updateSlice function,
	the function will still update the original mySlice slice, because
	when we create a slice in the Go, it actually creates two data structures;
	an underlying array that contains the specified elements and a slice that contains:
	a pointer to head (underlying array), capacity and length of slice.

	In updateSlice, Go is actually creating a copy of the data structure (slice), but that copy
	still points to the original array containing the elements ("Hi", "there" etc.)
	*/

	fmt.Println(mySlice);
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

/*
Value types in go: int, float, string, bool, structs
	- must use pointers to change these things in a function
Reference types: slices, maps, channels, pointers, functions
	- no need to worry about pointers with these
*/