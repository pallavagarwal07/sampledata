package a

import (
	"fmt"
	"time"
)

func Init() {
	fmt.Println("starting a")
	time.Sleep(1 * time.Second)
	fmt.Println("done a")
}
