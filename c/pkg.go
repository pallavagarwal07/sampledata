package c

import (
	"fmt"
	"time"
)

func Init() {
	fmt.Println("starting c")
	time.Sleep(1 * time.Second)
	fmt.Println("done c")
}
