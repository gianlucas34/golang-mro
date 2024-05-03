package main

import (
	"fmt"
	"time"

	"github.com/gianlucas34/golang-mro/internal"
)

func main() {
	fmt.Printf("Execution started...")

	start := time.Now()
	handler := internal.NewHandler()

	handler.Handle()

	end := time.Now()
	execTime := end.Sub(start)

	fmt.Printf("\nExecution completed in: %v", execTime)
}
