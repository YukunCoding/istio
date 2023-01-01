package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go RLock()
	go WLock()
	go Lock()
	time.Sleep(time.Second * 5)
}

// 读读不互斥
func RLock() {
	lock := sync.RWMutex{}
	for i := 1; i <= 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("RLock: ", i)
	}
}

// 写写互斥
func WLock() {
	lock := sync.RWMutex{}
	for i := 1; i <= 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("WLock: ", i)
	}
}

// 互斥锁
func Lock() {
	lock := sync.Mutex{}
	for i := 1; i <= 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("Lock: ", i)
	}
}
