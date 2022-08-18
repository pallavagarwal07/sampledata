package b

import (
	"fmt"
	"time"
)

func Init() {
	fmt.Println("starting b")
	time.Sleep(1 * time.Second)
	fmt.Println("done b")
}
