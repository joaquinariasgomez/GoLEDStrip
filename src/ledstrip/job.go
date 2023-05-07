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
	fmt.Println("Comenzando acción...", j.Action)
	time.Sleep(time.Second * 3)
	fmt.Println("Terminando acción")
}
