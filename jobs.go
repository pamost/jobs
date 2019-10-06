package jobs

import (
	"errors"
	"math/rand"
	"time"
)

type jobType func() error

// Jobs running different times
func job() jobType {
	return func() error {
		if rand.Float64() > 0.65 {
			return errors.New("err")
		}
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		return nil
	}
}

// Filling the slice with jobs
func sliceJobs(numJobs int) []jobType {
	var setJobs = []jobType{}
	for i := 0; i < numJobs; i++ {
		setJobs = append(setJobs, job())
	}
	return setJobs
}

// Parallel execution of N jobs
func handlerJobs(jobs []jobType, maxJobs int, maxErrors int) {
	var count int
	jobChannel := make(chan interface{}, len(jobs)) // Job execution channel
	errorsChannel := make(chan error, maxErrors)    // Channel with an error
	goroutines := make(chan struct{}, maxJobs)      // Semaphore to limit

	for _, j := range jobs {
		goroutines <- struct{}{}
		go func(j jobType) {
			if len(errorsChannel) < maxErrors {
				jobChannel <- j() // Writing completed jobs to the channel
				if j() != nil {
					errorsChannel <- j() // Writing jobs to the channel with an error
				}
			}
			<-goroutines
		}(j)
	}

	// Exit conditions from a function
	for range jobChannel {
		count++
		if count == len(jobs) || len(errorsChannel) == maxErrors {
			break
		}
	}
}
