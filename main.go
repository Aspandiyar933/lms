package main

import (
	handlers "command-line-arguments/Users/user/Documents/lms/handlers/hello.go"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.
	http.ListenAndServe(":9090", nil)
}