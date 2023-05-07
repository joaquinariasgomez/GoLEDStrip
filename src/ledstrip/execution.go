package ledstrip

import uuid "github.com/satori/go.uuid"

type execution struct {
	currentJob Job
	// Data about time of last execution, execution time, etc
}

func (e *execution) StartTask(a Action) {
	// Crea un job con la accion asignada y la ejecuta
	// Primero comprueba si hay algún job activo. Si es asi, manda una orden para pararlo
	// Luego, espera a que esté muerto el wg e inicia el nuevo job
	// No hay ningún job creado, así que toca crear uno
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

	e.currentJob = newJob
	e.currentJob.Start()
}

func (e *execution) stopAndLaunchJob(a Action) {
	e.currentJob.Stop()
	e.createAndLaunchJob(a)
}

var executionInstance *execution

func GetExecutionInstance() *execution {
	if executionInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if executionInstance == nil {
			executionInstance = &execution{}
		}
	}

	return executionInstance
}
