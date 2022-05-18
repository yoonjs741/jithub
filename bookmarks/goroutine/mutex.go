package bookmarks

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

func Add1Mutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Lock.Lock()
		Count += 1
		Lock.Unlock()
	}
}

func Add1Sem(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		sem <- struct{}{}
		Count += 1
		<-sem
	}
}

func Sub1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		Count -= 1
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

func Sub1Sem(wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		sem <- struct{}{}
		Count -= 1
		<-sem
	}
}

func main() {
	//DoSync()
	DoConcurrent()
	//DoConcurrentMutex()
	//DoSemaphore(1)
	//DoSemaphore(2)
	//DoSemaphore(10)
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
