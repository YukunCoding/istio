package main

import (
	"fmt"

	_ "github.com/istio/chapter11/a"
	_ "github.com/istio/chapter11/b"
)

func init() {
	fmt.Println("init main")
}

func main() {
	/*
		init函数的执行顺序
		main
			- a
				- b
			- b
		-------------
		init b
		init a
		init main
		good
	*/
	fmt.Println("good")
}
