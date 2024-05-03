package internal

import "time"

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Upload(index int) {
	time.Sleep(time.Second)
}
