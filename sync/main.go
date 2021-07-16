package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	var wf sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 5; i++ {
		wf.Add(1)
		go Deposit(i*100, &wf, &lock)
	}
	wf.Wait()
	fmt.Println(Balance(&lock))
}
