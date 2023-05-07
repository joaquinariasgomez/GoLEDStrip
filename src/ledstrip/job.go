package ledstrip

import (
	"fmt"
	"time"
)

type Job struct {
	ID     string
	Action Action
}

type Action struct {
	Type string `json:"type"`
	Mode string `json:"mode"`
}

func (j *Job) Start() {
	fmt.Printf("Comenzando job %v con acción %v\n", j.ID, j.Action)
	time.Sleep(time.Second * 3)
	fmt.Printf("Terminando job %v\n", j.ID)
}

func (j *Job) Stop() {
	fmt.Printf("Parando job %v\n", j.ID)
}
