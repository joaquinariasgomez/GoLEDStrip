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
	// If it is "set-brightness", it won't create new jobs.
	if a.Type == "set-brightness" {
		fmt.Println("Set brightness task")
		// Idea: crear un job all que no se le va a esperar ni nada para settear el brightness
	} else {
		// If there are no jobs, it will create and launch one.
		// Otherwise, it will stop the current job and launch another.
		if e.currentJob.ID == "" {
			e.createAndStartJob(a)
		} else {
			newJob := new(Job)
			newJob.Create(a)
			e.stopCurrentJobAndWaitForFinish()

			e.currentJob = *newJob
			e.currentJob.Start()
		}
	}
}

func (e *execution) StopTask() {

}

func (e *execution) createAndStartJob(a Action) {
	newJob := new(Job)
	newJob.Create(a)

	e.currentJob = *newJob
	e.currentJob.Start()
}

func (e *execution) stopCurrentJobAndWaitForFinish() {
	// Stop current execution if it isn't already
	if e.currentJob.status != "stopped" {
		e.currentJob.Stop()
	}
	// Wait for the current execution to finish
	e.currentJob.wg.Wait()
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
