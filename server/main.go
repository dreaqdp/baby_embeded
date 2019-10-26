package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Welcome to api"); err != nil {
		log.Printf("Error serving home. Error: %s", err.Error())

	}
	log.Print("conn to home")
}

func main() {
	http.HandleFunc("/", homePageEndpoint)
	log.Fatal(http.ListenAndServe(":2345", nil))
}
