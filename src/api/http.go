package api

import (
	"fmt"
	. "goledserver/src/constants"
	"log"
	"net/http"
)

// func setEndpoints() {
// 	mux.HandleFunc("/", HomePage)
// 	mux.HandleFunc("/action", GetAction).Methods("GET")
// 	mux.HandleFunc("/action", PostAction).Methods("POST")
// }

func HandleRequests() {
	fmt.Println("Hi! Welcome to Go LED Strip Service :)")
	fmt.Println("We are currently running in port", PORT)

	mux := http.NewServeMux()


	mux.HandleFunc("/randomFloat", func(w http.ResponseWriter,
		r *http.Request) {
		 fmt.Fprintln(w, 5.78)
		})
	// setEndpoints(mux)
	//mux.HandleFunc("/", HomePage)
	// mux.HandleFunc("/action", PostAction).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
