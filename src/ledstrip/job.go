package ledstrip

import (
	"fmt"
	"sync"
)

type Job struct {
	ID     string
	wg     sync.WaitGroup
	status string
	Action Action
}

type Action struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}

func (j *Job) Start() {
	defer j.wg.Done()
	j.status = "running"
	fmt.Printf("> Comenzando job %v con acción %v\n", j.ID, j.Action)
	TestAction(j.Action)
	fmt.Printf("Terminando job %v\n", j.ID)
}

func (j *Job) Stop() {
	j.status = "stopped"
	SetStopState()
	fmt.Printf("Parando job %v\n", j.ID)
}
