package shadowAgent

import (
	"context"
	"fmt"
	"log"
)

var jobID int64

// Enqueue used to set the job inside the queue
func Enqueue(ctx context.Context, job Job) error {
	if !ready {
		return fmt.Errorf("the worker is not ready yet")
	}
	if ctx != nil {
		job.Ctx = ctx
	}
	if job.Ctx == nil {
		job.Ctx = context.Background()
	}

	jobName := job.JobName
	if len(jobName) == 0 {
		return fmt.Errorf("invalid jobName")
	}

	if job.jobID == 0 {
		job.jobID = 1
	}
	job.jobID = jobID
	jobID++

	//write new data into jobs
	jobs := jm.pushJob(jobName, job)
	log.Printf("[shadowAgent] done queueing job %s, jobID: %d, totalTask: %d", jobName, job.jobID, len(jobs))
	return nil
}
