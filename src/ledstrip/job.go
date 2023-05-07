package ledstrip

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     string
	wg     sync.WaitGroup
	Action Action
}

type Action struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}

func (j *Job) Start() {
	defer j.wg.Done()
	fmt.Printf("> Comenzando job %v con acción %v\n", j.ID, j.Action)
	time.Sleep(time.Second * 5)
	fmt.Printf("Terminando job %v\n", j.ID)
}

func (j *Job) Stop() {
	fmt.Printf("Parando job %v\n", j.ID)
}
