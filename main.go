package main

import (
	"github.com/Aspandiyar933/lms/tree/main/handlers"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle('/', hh)
	http.ListenAndServe(":9090", sm)
}