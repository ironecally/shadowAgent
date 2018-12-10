package shadowAgent

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Job struct {
	JobName     string
	Ctx         context.Context
	HandlerFunc func(message []byte) error
	Message     []byte
	Action      string
	Sender      string

	jobID int64
}

type JobManager struct {
	mtx      sync.RWMutex
	jobQueue []Job
}

var jm *JobManager

func initJobManager() *JobManager {
	jm = &JobManager{
		mtx:      sync.RWMutex{},
		jobQueue: []Job{},
	}

	return jm
}

func (jm *JobManager) pushJob(jobName string, job Job) []Job {
	jm.mtx.RLock()
	defer jm.mtx.RUnlock()
	log.Println("pushing job", jobName)
	jm.jobQueue = append(jm.jobQueue, job)
	return jm.jobQueue
}

func (jm *JobManager) popJob(jobName string) (Job, error) {
	jm.mtx.RLock()
	defer jm.mtx.RUnlock()
	//check if there's no data
	if len(jm.jobQueue) == 0 {
		return Job{}, fmt.Errorf("no data")
	}
	job := jm.jobQueue[0]

	log.Println("popping job", jobName)
	//check if there's no other data
	if len(jm.jobQueue) > 1 {
		jm.jobQueue = jm.jobQueue[1:]
	} else {
		jm.clearJob(jobName)
	}

	return job, nil
}

func (jm *JobManager) clearJob(jobName string) {
	jm.mtx.RLock()
	defer jm.mtx.RUnlock()

	jm.jobQueue = []Job{}
	return
}
