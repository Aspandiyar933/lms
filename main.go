package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r*http.Request){
		log.Println("Hello, World")
		d, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, "Hello, %s", d)
	})

	http.HandleFunc("/aspan", func(http.ResponseWriter, *http.Request){
		log.Println("Hello, Aspan")
	})

	http.ListenAndServe(":9090", nil)
}