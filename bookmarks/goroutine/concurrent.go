package main

import (
	"fmt"
	"sync"
)

var (
	Count int         = 0
	Lock  *sync.Mutex = &sync.Mutex{}
)

func Add1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Count += 1
	}
}

func Sub1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Count -= 1
	}
}

func DoConcurrent() {
	for i := 0; i < 5; i++ {
		Count = 0
		wg := &sync.WaitGroup{}
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go Add1(wg)
			go Sub1(wg)
		}
		wg.Wait()
		fmt.Println("Concurrent goroutines Result(Desired 0):", Count, "")
	}

	fmt.Println("=====================================")
}

func main() {
	// DoConcurrent()
	// DoSync()
	for i := 0; i < 5; i++ {
		Count = 0
		wg := &sync.WaitGroup{}
		wg.Add(20)
		for i := 0; i < 10; i++ {
			go Add1(wg)
			go Sub1(wg)
		}
		wg.Wait()
		fmt.Println("Concurrent goroutines Result(Desired 0):", Count, "")
	}

	fmt.Println("=====================================")
}

func DoSync() {
	for i := 0; i < 5; i++ {
		Count = 0
		wg := &sync.WaitGroup{}
		wg.Add(20)
		for i := 0; i < 10; i++ {
			Add1(wg)
			Sub1(wg)
		}
		wg.Wait()
		fmt.Println("Sync Result(Desired 0):", Count, "")
	}

}
