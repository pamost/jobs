package jobs

import (
	"testing"
)

func TestHandlerJobs(t *testing.T) {
	table := []struct {
		jobs      int
		maxErrors int
		maxJobs   int
	}{
		{5, 1, 1},
		{10, 3, 1},
		{50, 5, 5},
	}
	for _, row := range table {
		handlerJobs(sliceJobs(row.jobs), row.maxJobs, row.maxErrors)
	}
}
