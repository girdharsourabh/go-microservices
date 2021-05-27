package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		resp := "Hello world"
		rw.Write([]byte(resp))
	})

	http.ListenAndServe(":9000", nil)
}
