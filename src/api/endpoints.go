package api

import (
	"encoding/json"
	"fmt"
	"goledserver/src/ledstrip"
	"io"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
}

func Action(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostAction(w, r)
	} else {
		GetAction(w, r)
	}
}

func GetAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recibida acción GET:", r)

	//ledstrip.ExampleWipe()

	w.WriteHeader(http.StatusOK)
	/*w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)*/
}

func PostAction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		resp := make(map[string]string)
		resp["message"] = "Error: Request Body needed in POST Action operation"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}
	if bodyBytes, err := io.ReadAll(r.Body); err == nil {
		var action ledstrip.Action
		if err = json.Unmarshal(bodyBytes, &action); err != nil {
			badReqResp := make(map[string]string)
			badReqResp["message"] = "Error unmarshalling request body"
			jsonBadReqResp, _ := json.Marshal(badReqResp)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonBadReqResp)
			return
		}

		execution := ledstrip.GetExecutionInstance()
		go execution.StartTask(action)

		w.WriteHeader(http.StatusCreated)
		w.Write(bodyBytes)
	} else {
		badReqResp := make(map[string]string)
		badReqResp["message"] = "Error reading request body"
		jsonBadReqResp, _ := json.Marshal(badReqResp)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonBadReqResp)
	}
}
