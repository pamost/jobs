package jobs

import (
	"errors"
	"math/rand"
	"time"
)

type jobType func() error

// job - для заданий выполняющихся разное время
func job() jobType {
	return func() error {
		if rand.Float64() > 0.65 {
			return errors.New("err")
		}
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		return nil
	}
}

//sliceJobs - наполнение слайса заданиями
func sliceJobs(numJobs int) []jobType {
	var setJobs = []jobType{}
	for i := 0; i < numJobs; i++ {
		setJobs = append(setJobs, job())
	}
	return setJobs
}

// handlerJobs - функция параллельного выполнение N заданий
func handlerJobs(jobs []jobType, maxJobs int, maxErrors int) {
	var count int
	jobСhannel := make(chan interface{}, len(jobs)) // Канал выполенния заданий
	errorsСhannel := make(chan error, maxErrors)    // Канал заданий с ошибкой
	goroutines := make(chan struct{}, maxJobs)      // Канал ограничения горутин

	for _, j := range jobs {
		func(j jobType) {
			goroutines <- struct{}{}
			if len(errorsСhannel) < maxErrors {
				jobСhannel <- j() // Запись в канал выполненных заданий
				if job() != nil {
					errorsСhannel <- j() // Запись в канал заданий с ошибкой
				}
			}
			<-goroutines
		}(j)
	}

	// Выход из функции
	for range jobСhannel {
		count++
		if count == len(jobs) || len(errorsСhannel) == maxErrors {
			break
		}
	}
}
