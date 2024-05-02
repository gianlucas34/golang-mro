package internal

import "sync"

type Task interface {
	Process()
}

type WorkerPool struct {
	Tasks       []Task
	Concurrency int
	Channel     chan Task
	Wg          sync.WaitGroup
}

func NewWorkerPool(tasks []Task, concurrency int) *WorkerPool {
	return &WorkerPool{
		Tasks:       tasks,
		Concurrency: concurrency,
	}
}

func (wp *WorkerPool) worker() {
	for task := range wp.Channel {
		task.Process()
	}

	defer wp.Wg.Done()
}

func (wp *WorkerPool) Run() {
	wp.Channel = make(chan Task)

	wp.Wg.Add(wp.Concurrency)

	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	for _, task := range wp.Tasks {
		wp.Channel <- task
	}

	close(wp.Channel)

	wp.Wg.Wait()
}
