package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)



func main()  {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "hello %s", d)
	})

	http.ListenAndServe(":8080", nil)
}