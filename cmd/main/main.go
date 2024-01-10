package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiyoshi-87/RetireMate-GO-backend/pkg/routes"
)

func main() {
	r := routes.Router()

	fmt.Println("Starting server on the port 4000...")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}
