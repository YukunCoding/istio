package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	src, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer src.Close()
	copy(src, f)
}

// copy file
func copy(src, dest *os.File) {
	buffer := make([]byte, 64)
	for len := -1; len != 0; {
		len, _ = src.Read(buffer)
		fmt.Println(string(buffer[:len]))
		dest.WriteString(string(buffer[:len]))
	}
}
