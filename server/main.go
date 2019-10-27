package main

import (
	"log"
	"net/http"
)

var m *ArduinoManager

func main() {
	m = NewArduinoManager()
	go m.run()
	go m.attendReq()
	http.HandleFunc("/", homePageEndpoint)
	http.HandleFunc("/forward", forwardEndpoint)
	http.HandleFunc("/back", backEndpoint)
	http.HandleFunc("/left", leftEndpoint)
	http.HandleFunc("/right", rightEndpoint)
	http.HandleFunc("/stop", stopEndpoint)
	http.HandleFunc("/play", audioEndpoint)
	log.Fatal(http.ListenAndServe(":2345", nil))
}
