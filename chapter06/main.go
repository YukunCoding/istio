package main

import (
	"sync"
	"time"
)

func main() {
	safeWrite()
	time.Sleep(time.Second * 4)
}

// 普通map的不安全写操作
func unsafeWrite() {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		go func(i int) {
			m[0] = i
		}(i)
	}
}

type SafeMap struct {
	data  map[int]int
	mutex sync.Mutex
}

func (s *SafeMap) read(idx int) (int, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	result, ok := s.data[idx]
	return result, ok
}

func (s *SafeMap) write(idx, item int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[idx] = item
}

func safeWrite() {
	s := SafeMap{
		data:  map[int]int{},
		mutex: sync.Mutex{},
	}
	for i := 0; i < 200; i++ {
		go func() {
			s.write(1, 1)
		}()
	}
}
