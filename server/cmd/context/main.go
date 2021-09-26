package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct {
}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	ctx, cancelFunc := context.WithTimeout(c, 5*time.Second)
	defer cancelFunc()
	mainTask(ctx)
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))

	//c1, cancelFunc := context.WithTimeout(c, time.Second*2)
	//defer cancelFunc()

	go smallTask(context.Background(), "task1", 4*time.Second)
	smallTask(c, "task2", 1*time.Second)

}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s started wit param %q \n", name, c.Value(paramKey{}))
	select {
	case <-time.After(d):
		fmt.Printf("%s done\n", name)
	case <-c.Done():
		fmt.Printf("%s cancelled\n", name)

	}
}
