package api

import (
	"fmt"
	"log"
	"net/http"
	. "goledserver/src/constants"
)

func setEndpoints() {
	http.HandleFunc("/", HomePage)
}

func HandleRequests() {
	fmt.Println("Hi! Welcome to Go LED Strip Service :)")
	fmt.Println("We are currently running in port",PORT)

	setEndpoints()
	
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}