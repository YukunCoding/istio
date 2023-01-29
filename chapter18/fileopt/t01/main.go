package main

import(
	"os"
	"fmt"
)


func main() {
	file, err := os.Open("/home/lighthouse/a.md") // for read access.
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 100)
	count, error := file.Read(data)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf(string(data[:count]))
}


