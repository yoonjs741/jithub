package main

import (
	"fmt"
	"sync"
)

func Add1Mutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Lock.Lock()
		Count += 1
		Lock.Unlock()
	}
}

func Sub1Mutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Lock.Lock()
		Count -= 1
		Lock.Unlock()
	}
}

func DoConcurrentMutex() {
	for i := 0; i < 5; i++ {
		Count = 0
		wg := &sync.WaitGroup{}
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go Add1Mutex(wg)
			go Sub1Mutex(wg)
		}
		wg.Wait()
		fmt.Println("Concurrent goroutines + Mutex Result(Desired 0):", Count, "")
	}

	fmt.Println("=====================================")
}
