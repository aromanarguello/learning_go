package main

import (
	"io"
	"net/http"
	"fmt"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}
	
	io.Copy(lw, resp.Body)
	// fmt.Println(lw)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes written:", len(bs))
	return len(bs), nil
}