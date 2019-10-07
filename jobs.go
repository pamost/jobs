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
func makeSliceJobs(numJobs int) []jobType {
	var setJobs = []jobType{}
	for i := 0; i < numJobs; i++ {
		setJobs = append(setJobs, job())
	}
	return setJobs
}

// Parallel execution of N jobs
func handlerJobs(sliceJobs []jobType, maxJobs int, maxErrors int) int {
	var countJobs, countErrors int
	jobChannel := make(chan interface{}, len(sliceJobs)) // Job execution channel
	errorsChannel := make(chan error, maxErrors)         // Channel with an error
	goroutines := make(chan struct{}, maxJobs)           // Semaphore to limit

	for _, j := range sliceJobs {
		go func(j jobType) {
			goroutines <- struct{}{}
			if len(errorsChannel) < maxErrors {
				process(j, jobChannel, errorsChannel)
			}
			<-goroutines
		}(j)
	}

	for range jobChannel {
		countJobs++
		if countJobs == len(sliceJobs) || len(errorsChannel) == maxErrors {
			countErrors = len(errorsChannel)
			break
		}
	}
	return countErrors
}

// Job processing
func process(j jobType, jobChannel chan<- interface{}, errorsChannel chan<- error) {

	result := j()
	jobChannel <- result // Writing completed jobs to the channel

	if result != nil {
		errorsChannel <- result // Writing jobs to the channel with an error
	}
}
