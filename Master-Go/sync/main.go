package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	// executes this before finishing the function execution -1
	defer wg.Done()
	lock.Lock()
	// protects this modificaiton with a lock
	balance += amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	// protects the value from being modified while reading with Rlock
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	// Wait group to wait for gouroutines to finish
	var wg sync.WaitGroup
	// RW lock: gives us the opportunity to block
	// accesing and modifying the variable when trying
	// to write with lock.Lock() and lock.Unlock()
	// While at the same time protects the value `balance`
	// from being modified while reading with lock.RLock() and
	// freeing with lock.RUnlock()
	var lock sync.RWMutex
	// 5 goroutines
	for i := 1; i <= 5; i++ {
		// Wait Group counter +1
		wg.Add(1)
		// gouritne and we pass wg and lock
		go Deposit(i*100, &wg, &lock)
	}
	// Waits until the wg counter is 0
	wg.Wait()
	fmt.Println(Balance(&lock))

}
