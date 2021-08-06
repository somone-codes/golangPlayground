package main

import "fmt"

func main() {
	//select switch (channel)
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)
}
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
	quit <- true
}

// receive channel
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v, ok := <-even:
			if !ok {
				fmt.Println("Channel closed!")
			} else {
				fmt.Println("the value received from the even channel:", v)
			}
		case v, ok := <-odd: // optionally using comma-ok idiom to check if value received or not
			if !ok {
				fmt.Println("Channel closed!")
			} else {
				fmt.Println("the value received from the even channel:", v)
			}
		case <-quit:
			return
		}
	}
}
