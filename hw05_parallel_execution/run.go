package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errorsCount int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	tasksCount := len(tasks)
	wg.Add(n)

	ch := make(chan Task, tasksCount)

	if n > tasksCount {
		n = tasksCount
	}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			isContinue := true
			for task := range ch {
				err := task()

				mu.Lock()
				if err != nil {
					errorsCount++
				}

				if errorsCount >= m {
					isContinue = false
				}

				mu.Unlock()

				if m <= 0 {
					continue
				}

				if isContinue == false {
					return
				}
			}
		}()
	}

	for _, task := range tasks {
		ch <- task
	}
	close(ch)

	wg.Wait()

	if m > 0 && errorsCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
