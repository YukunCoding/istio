package main

import (
	"fmt"
	"sync"
)

// 重命名数组
type SliceNum []int

// 构造方法，返回一个int数组
func NewSlice() SliceNum {
	return make(SliceNum, 0)

}

// 向数组中添加元素
func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	return s
}

func main() {
	var once sync.Once
	s := NewSlice()
	// 看源代码理解once的行为
	once.Do(func() {
		s.Add(89)
	})
	once.Do(func() {
		s.Add(99)
	})
	once.Do(func() {
		s.Add(11)
	})
	fmt.Println(s)
}
