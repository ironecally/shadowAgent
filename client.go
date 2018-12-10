package shadowAgent

import (
	"log"
)

var jobBuffer chan Job
var errBuffer chan errType

var ready bool

//Option shall be only used for preparing new queue
type Option struct {
	TotalWorker int
}

//NewQueue is to init the queueing system
func NewQueue(opt *Option) {
	if opt.TotalWorker <= 0 {
		opt.TotalWorker = 5
	}

	jm = initJobManager()
	jobBuffer = make(chan Job, opt.TotalWorker)
	errBuffer = make(chan errType, opt.TotalWorker)
	ready = true
	log.Printf("[shadowAgent] running with opt %+v", opt)
	for {
		//always check and try to empty the job queue
		go checkAllQueue()

		select {
		case job := <-jobBuffer:
			execute(job)
		case errFound := <-errBuffer:
			errHandler(errFound)
		}
	}

}

func checkAllQueue() {
	for {
		job, err := jm.popJob("testing")
		if err != nil {
			// time.Sleep(1 * time.Second)
			continue
		}

		if job.jobID == 0 {
			continue
		}
		jobBuffer <- job
	}
}
