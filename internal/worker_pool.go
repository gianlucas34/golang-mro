package internal

type Job struct {
	Adapter Adapter
	Index   int
}

type WorkerPool struct {
	Jobs        []Job
	Concurrency int
	JobChan     chan Job
	DoneChan    chan struct{}
}

func NewWorkerPool(jobs []Job, concurrency int) *WorkerPool {
	return &WorkerPool{
		Jobs:        jobs,
		Concurrency: concurrency,
		JobChan:     make(chan Job, len(jobs)),
		DoneChan:    make(chan struct{}),
	}
}

func (wp *WorkerPool) worker() {
	for job := range wp.JobChan {
		job.Adapter.Upload(job.Index)
	}

	wp.DoneChan <- struct{}{}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	for _, job := range wp.Jobs {
		wp.JobChan <- job
	}

	close(wp.JobChan)

	for i := 0; i < wp.Concurrency; i++ {
		<-wp.DoneChan
	}

	close(wp.DoneChan)
}
