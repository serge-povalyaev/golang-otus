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
	wg.Add(n)

	ch := make(chan Task, len(tasks))

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for task := range ch {
				err := task()
				if err != nil {
					mu.Lock()
					errorsCount++
					mu.Unlock()
				}

				if m <= 0 {
					continue
				}

				if errorsCount >= m {
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
