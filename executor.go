package shadowAgent

func execute(job Job) {
	err := job.HandlerFunc(job.Message)
	if err != nil {
		errBuffer <- errType{err, job}
	}

	return
}
