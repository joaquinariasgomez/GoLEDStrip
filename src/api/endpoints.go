package api

import (
	"fmt"
	"net/http"
	"goledserver/src/ledstrip"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
	ledstrip.StartLedStrip()
}