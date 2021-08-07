package main

import (
	"context"
	"fmt"
)

func main() {
	context, cancel := context.WithCancel(context.Background())

	defer cancel() // this will be executed when the for loop breaks and this will in turn causes generator to exit

	for i := range generator(context) {
		if i < 10 {
			fmt.Println(i)
			continue
		}
		break
	}

}

//given a context this function returns a channel from which number can be read from,
//generates number starting from 0 until context is done
func generator(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 0

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Got cancel signal.")
				return
			case ch <- n:
				fmt.Println("Pushing number to channel")
				n++
			}
		}
	}()
	return ch
}
