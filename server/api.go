package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageEndpoint(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Welcome to api"); err != nil {
		log.Printf("Error serving home. Error: %s", err.Error())

	}
	log.Print("conn to home")
}
func audioEndpoint(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print("Error parsing audio post")
	} else {
		//	url := r.Form.Get("yt")
	}

}

func stopEndpoint(w http.ResponseWriter, r *http.Request) {
	m.toSend <- '0'
}
func forwardEndpoint(w http.ResponseWriter, r *http.Request) {
	m.toSend <- '1'
}
func leftEndpoint(w http.ResponseWriter, r *http.Request) {
	m.toSend <- '2'
}
func rightEndpoint(w http.ResponseWriter, r *http.Request) {
	m.toSend <- '3'
}
func backEndpoint(w http.ResponseWriter, r *http.Request) {
	m.toSend <- '4'
}
