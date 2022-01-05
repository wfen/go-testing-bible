package main

import (
	"fmt"
	"log"
	"net/http"
)

// HomePage - our homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Setup - sets up our routes
func Setup() {
	http.HandleFunc("/", HomePage)
}

func main() {
	Setup()
	log.Fatal(http.ListenAndServe(":10000", nil))
}
