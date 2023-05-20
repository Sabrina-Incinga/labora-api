package main

import "sync"

func Incrementx100(counter *int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		*counter++
	}
	wg.Done()
}
