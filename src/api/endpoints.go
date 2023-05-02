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

func WipeBlue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Golpeado endpoint blue!")
	ledstrip.ExampleWipe()
}

func WipeRed(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Golpeado endpoint red!")
	ledstrip.ExamplePulsate()
}
