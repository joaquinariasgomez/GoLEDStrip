package api

import (
	"fmt"
	"goledserver/src/ledstrip"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	ledstrip.PrintDeviceStatus()
}

func GetAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibida acción GET:", r)
	ledstrip.ExampleWipe()
}

func PostAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibida acción POST:", r)
	ledstrip.ExampleWipe()
}
