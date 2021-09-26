package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.Tick(2 * time.Second):
		fmt.Println("2 second over", time.Now().Second())
	case <-time.After(5 * time.Second):
		fmt.Println("5 second over timeover", time.Now().Second())
		return

	}
}
