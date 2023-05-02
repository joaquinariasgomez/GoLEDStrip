package api

import (
	"fmt"
	"goledserver/src/ledstrip"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
	// ledstrip.ExampleStartAndWipe()
	ledstrip.PrintDeviceStatus()
}

func WipeBlue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Golpeado endpoint blue!")
	// ledstrip.WipeBlue()
	ledstrip.ExampleWipe()
}

func WipeRed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Golpeado endpoint red!")
	// ledstrip.WipeRed()
	// ledstrip.ShutdownDevice()
}
