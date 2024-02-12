package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
		log.Println("Hello, World")
	})

	http.HandleFunc("/aspan", func(http.ResponseWriter, *http.Request){
		log.Println("Hello, Aspan")
	})

	http.ListenAndServe(":9090", nil)
}