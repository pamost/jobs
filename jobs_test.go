package jobs

import (
	"testing"
)

func TestHandlerJobsErrors(t *testing.T) {
	table := []struct {
		jobs      int
		maxJobs   int
		maxErrors int
	}{
		{5, 2, 1},
		{10, 5, 9},
		{50, 6, 5},
	}
	for _, row := range table {
		countErr := handlerJobs(makeSliceJobs(row.jobs), row.maxJobs, row.maxErrors)
		//fmt.Printf("Func with errors: %d == %d \n", countErr, row.maxErrors)
		if countErr > row.maxErrors {
			t.Errorf("functions with errors %d, specified limit %d", countErr, row.maxErrors)
		}
	}
}
