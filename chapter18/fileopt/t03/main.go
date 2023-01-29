package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "kate")
	os.Setenv("BURROW", "/usr/gopher")
	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
}
