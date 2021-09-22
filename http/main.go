package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	*/

	// Same thing as code above
	// io.Copy(os.Stdout, resp.Body)

	// With custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// logWriter type now implements Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes:", len(bs))
	
	return len(bs), nil
}