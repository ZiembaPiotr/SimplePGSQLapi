package main

import (
	"github.com/mux/router/router"
	"log"
	"net/http"
)

func main() {
	r := router.CreateRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
