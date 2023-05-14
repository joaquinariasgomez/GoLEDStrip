package api

import (
	"fmt"
	. "goledserver/src/constants"
	"goledserver/src/ledstrip"
	"log"
	"net/http"
)

func setEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/", HomePage)
	mux.HandleFunc("/action", Action)
}

func runStartupSequence() {
	execution := ledstrip.GetExecutionInstance()
	action := ledstrip.Action{
		Type: ledstrip.Startup,
		Command: ledstrip.Command{
			Instruction: "startup-animation",
			Args:        nil,
		},
	}

	go execution.StartTask(action)
}

func HandleRequests() {
	fmt.Println("Hi! Welcome to Go LED Strip Service :)")
	fmt.Println("We are currently running in port", PORT)

	mux := http.NewServeMux()
	setEndpoints(mux)

	runStartupSequence()
	log.Fatal(http.ListenAndServe(":"+PORT, mux))
}
