package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Welcome to api"); err != nil {
		log.Printf("Error serving home. Error: %s", err.Error())

	}
	log.Print("conn to home")
}

func main() {
	base := ""
	http.HandleFunc("/", homePageEndpoint)
	log.Fatal(http.ListenAndServeTLS(":2345", path.Join(base, "server.crt"), path.Join(base, "./server.key"), nil))
}
