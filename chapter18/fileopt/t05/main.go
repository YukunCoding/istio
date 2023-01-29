package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("gid: %d, uid: %d.\n", os.Getuid(), os.Getgid())
	fmt.Printf("get pagesize: %d.\n", os.Getpagesize())
}
