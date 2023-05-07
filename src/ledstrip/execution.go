package ledstrip

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

var executionLock = &sync.Mutex{}

type execution struct {
	currentJob Job
	// Data about time of last execution, execution time, etc
}

func (e *execution) StartTask(a Action) {
	// If there are no jobs, it will create and launch one
	// Otherwise, it will stop the current job and launch another
	if e.currentJob.ID == "" {
		e.createAndLaunchJob(a)
	} else {
		e.stopAndLaunchJob(a)
	}
}

func (e *execution) StopTask() {

}

func (e *execution) createAndLaunchJob(a Action) {
	var newJob Job
	newJob.ID = uuid.NewV4().String()
	newJob.Action = a
	newJob.wg = sync.WaitGroup{}
	newJob.wg.Add(1)

	e.currentJob = newJob
	e.currentJob.Start()
}

func (e *execution) stopAndLaunchJob(a Action) {
	// Stop current execution if it isn't already
	if e.currentJob.status != "stopped" {
		e.currentJob.Stop()
	}
	// Wait for the current execution to finish
	e.currentJob.wg.Wait()
	e.createAndLaunchJob(a)
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
