package main

import (
	"fmt"
	"sync"
)

var counter int
var counterWithMutex int

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex


	wg.Add(2)
	go Incrementx100(&counter, &wg)
	go Incrementx100(&counter, &wg)


	wg.Add(2)
	go Incrementx100WithMutex(&counterWithMutex, &wg, &mu)
	go Incrementx100WithMutex(&counterWithMutex, &wg, &mu)

	wg.Wait()
	fmt.Println("Final value of counter is: ", counter)
	fmt.Println("Final value of counterWithMutex is: ", counterWithMutex)

}