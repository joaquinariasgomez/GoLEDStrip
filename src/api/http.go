package api

import (
	"fmt"
	. "goledserver/src/constants"
	"log"
	"net/http"
)

func setEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/action", Action)
}

func HandleRequests() {
	fmt.Println("Hi! Welcome to Go LED Strip Service :)")
	fmt.Println("We are currently running in port", PORT)

	mux := http.NewServeMux()
	setEndpoints(mux)

	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
