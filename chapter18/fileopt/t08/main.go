package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("start")
	f, err := os.OpenFile("./a.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

