package main

import "fmt"

func main() {
	var myArr = []string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < len(myArr); i++ {
		if myArr[i] == "stupid" {
			myArr[i] = "smart"
		}
		if myArr[i] == "weak" {
			myArr[i] = "strong"
		}
	}

	for _, char := range myArr {
		fmt.Printf("%s ", char)
	}
}
