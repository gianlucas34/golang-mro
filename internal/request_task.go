package internal

import "time"

type RequestTask struct{}

func NewRequestTask() *RequestTask {
	return &RequestTask{}
}

func (rt *RequestTask) Process() {
	time.Sleep(time.Second * 4)
}
