package gobyexample

import (
	"fmt"
	"time"
)

// In this example we'll look at how to implement a worker pool using
// goroutines and channels.

// Here's the worker, of whih we'll run several concurrent instances.
// These workers will receive work on the `jobs` channel and send the
// results to the `results` channel. We'll sleep a channel per job to
// simulate an expensive task.
func workerInPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// WorkerPoolsDemo - demonstrates running a worker pool with
// goroutines and channels.
func WorkerPoolsDemo() {
	// In order to use our pool of workers, we need to send them
	// work and collect their results. We make two channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start up 3 workers. Initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go workerInPool(w, jobs, results)
	}

	// Send 5 jobs and close that channel to indicate that's all the work we have.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results of all the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
