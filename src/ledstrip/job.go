package ledstrip

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type Job struct {
	ID     string
	wg     *sync.WaitGroup
	status string
	Action Action
}

type Action struct {
	Type    ActionTypeEnum `json:"type"`
	Command Command        `json:"command"`
}

type Command struct {
	Instruction InstructionTypeEnum `json:"instruction"`
	Args        []string            `json:"args"`
}

type ActionTypeEnum string

const (
	Startup       ActionTypeEnum = "startup"
	ChangeMode    ActionTypeEnum = "change-mode"
	SetColor      ActionTypeEnum = "set-color"
	SetBrightness ActionTypeEnum = "set-brightness"
)

type InstructionTypeEnum string

const (
	OfficeLights InstructionTypeEnum = "office-lights"
	StaticColor  InstructionTypeEnum = "static-color"
	Increase     InstructionTypeEnum = "increase"
	Decrease     InstructionTypeEnum = "decrease"
)

func (j *Job) Create(a Action) {
	j.status = "waiting"
	j.ID = uuid.NewV4().String()
	j.Action = a
	j.wg = &sync.WaitGroup{}
	j.wg.Add(1)

	//fmt.Printf("> Creando job %v en estado %s con acción %v\n", j.ID, j.status, j.Action)
}

func (j *Job) Start() {
	defer j.wg.Done()
	j.status = "running"

	fmt.Printf("> Comenzando job %v en estado %s con acción %v\n", j.ID, j.status, j.Action)
	StartAction(j.Action)
	fmt.Printf("< Terminando job %v\n", j.ID)
}

func (j *Job) Stop() {
	j.status = "stopped"
	SetStopState()
	fmt.Printf("Parando job %v\n", j.ID)
}
