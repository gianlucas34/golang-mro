package main

import (
	"fmt"
	"time"

	"github.com/gianlucas34/golang-mro/internal"
)

func main() {
	fmt.Printf("Execução iniciada...")

	start := time.Now()
	tasks := make([]internal.Task, 0)

	for i := 0; i < 800; i++ {
		tasks = append(tasks, internal.NewRequestTask())
	}

	wp := internal.NewWorkerPool(tasks, 20)

	wp.Run()

	end := time.Now()
	execTime := end.Sub(start)

	fmt.Printf("\nExecução finalizada em: %v", execTime)
}
