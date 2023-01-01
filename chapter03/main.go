package main

import (
	"fmt"
	"sync"
	"time"
)

// 类似于JAVA当中的阻塞队列
type Queue struct {
	data []string
	cond *sync.Cond
}

func main() {
	queue := Queue{
		data: []string{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
	go func() {
		// 生产者线程
		for {
			queue.Enqueue("kate")
			time.Sleep(time.Second)
		}
	}()
	for {
		// 消费者线程
		queue.Dequeue()
	}
}

// 进入队列
func (q *Queue) Enqueue(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.data = append(q.data, item)
	fmt.Printf("add item %s\n", item)
	q.cond.Broadcast()
}

// 从队列中获取元素
func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.data) == 0 {
		fmt.Printf("no data, please wait a moment!\n")
		q.cond.Wait()
	}
	result := q.data[0]
	q.data = q.data[1:]
	return result
}
