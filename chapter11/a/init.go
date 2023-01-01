package a

import (
	"fmt"

	_ "github.com/istio/chapter11/b"
)

func init() {
	fmt.Println("init a")
}
