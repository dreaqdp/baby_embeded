package main

import (
	"encoding/json"
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
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	url := struct {
		Url string `json:"yt"`
	}{}

	if err := d.Decode(&url); err != nil {
		log.Print("Unrecognized json")
		return
	}

	log.Print("Playing ", url.Url)
	m.music <- url.Url

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
