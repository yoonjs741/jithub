package main

import (
	"fmt"
	"sync"
)

func Add1Sem(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		sem <- struct{}{}
		Count += 1
		<-sem
	}
}

func Sub1Sem(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		sem <- struct{}{}
		Count -= 1
		<-sem
	}
}

func DoSemaphore(maxConcurrent int64) {
	for i := 0; i < 5; i++ {
		Count = 0
		wg := &sync.WaitGroup{}
		sem := make(chan struct{}, maxConcurrent)
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go Add1Sem(wg, sem)
			go Sub1Sem(wg, sem)
		}
		wg.Wait()
		fmt.Println("Concurrent goroutines + Semaphore", maxConcurrent, "Result(Desired 0):", Count, "")
	}

	fmt.Println("=====================================")
}
