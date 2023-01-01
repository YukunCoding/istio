package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := WorkQueue{
		capacity: 5,
		data:     []string{},
		cond:     sync.NewCond(&sync.Mutex{}),
	}
	for i := 0; i <= 5; i++ {
		queue.Enqueue("kate")
	}

	time.Sleep(time.Second)
	queue.Dequeue()
}

// 定义一个有界工作队列
type WorkQueue struct {
	capacity int
	data     []string
	cond     *sync.Cond
}

// 判断队列是否为空
func (wq *WorkQueue) IsEmpty() (res bool) {
	res = len(wq.data) == 0
	return res
}

// 判断队列是否已满
func (wq *WorkQueue) IsFull() (res bool) {
	res = len(wq.data) == wq.capacity
	return res
}

// 添加任务到工作队列
func (wq *WorkQueue) Enqueue(item string) {
	wq.cond.L.Lock()
	defer wq.cond.L.Unlock()
	// // 如果队列已满，则阻塞等待
	// for wq.IsFull() {
	// 	fmt.Println("Full wait...")
	// 	wq.cond.Wait()
	// }
	wq.data = append(wq.data, item)
	wq.cond.Broadcast()
}

// 获取任务进行执行
func (wq *WorkQueue) Dequeue() string {
	wq.cond.L.Lock()
	defer wq.cond.L.Unlock()
	// 如果队列为空，则阻塞等待
	for wq.IsEmpty() {
		fmt.Println("Empty wait...")
		wq.cond.Wait()
	}
	result := wq.data[0]
	wq.data = wq.data[1:]
	// wq.fullCond.Broadcast()
	return result
}
