package d

import (
	"fmt"
	"time"
)

func Init() {
	fmt.Println("starting d")
	time.Sleep(1 * time.Second)
	fmt.Println("done d")
}
