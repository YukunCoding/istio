package main

import (
	"fmt"
	"os"
	"log"
	"time"
)

func main() {
	fs, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(fs.Name()) // clean up
	if _, err := fs.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := fs.Close(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 5)
}
	
