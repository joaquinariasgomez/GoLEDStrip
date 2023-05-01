package api

import (
	"fmt"
	"log"
	"net/http"
	. "goledserver/src/constants"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() {
	fmt.Println("Hi! Welcome to Go LED Strip Service :)")
	fmt.Println("We are currently running in port",PORT)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}