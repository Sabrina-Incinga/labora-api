package main

import "sync"

func Incrementx100WithMutex(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	for i := 0; i < 100; i++ {
		*counter++
	}
	mu.Unlock()
	wg.Done()
}
