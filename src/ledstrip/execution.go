package ledstrip

import (
	"fmt"
	"sync"
)

var executionLock = &sync.Mutex{}

type execution struct {
	currentJob Job
	// Data about time of last execution, execution time, etc
}

func (e *execution) StartTask(a Action) {
	// First, check what action type is being executed.
	// If it is "set-brightness", it won't create any jobs.
	// The flow for this action will be different
	if a.Type == SetBrightness {
		fmt.Println("Set brightness task")
		StartAction(a)
		// functions.BrightnessAction(a) -> Que dentro de functions se interprete la acción y se hable directamente con device
		// Idea: crear un job al que no se le va a esperar ni nada para settear el brightness
	} else {
		// We will always create a new job.
		// If there are no jobs in the execution, we will assign it and launch it
		// Otherwise, it will stop the current job and launch another.
		newJob := new(Job)
		newJob.Create(a)
		if e.currentJob.ID == "" {
			e.currentJob = *newJob
			e.currentJob.Start()
		} else {
			// Stop current execution if it isn't already
			if e.currentJob.status != "stopped" {
				e.currentJob.Stop()
			}
			// Wait for the current execution to finish
			e.currentJob.wg.Wait()

			e.currentJob = *newJob
			e.currentJob.Start()
		}
	}
}

func (e *execution) StopTask() {

}

var executionInstance *execution

func GetExecutionInstance() *execution {
	if executionInstance == nil {
		executionLock.Lock()
		defer executionLock.Unlock()

		if executionInstance == nil {
			executionInstance = &execution{}
		}
	}

	return executionInstance
}
