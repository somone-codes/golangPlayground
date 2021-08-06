package main

import "fmt"

func ping(writeChannel chan<- string) { // writeChannel is a read only channel here
	writeChannel <- "ping"
}

func pong(readChannel <-chan string) { // pings is a write only channel here and pongs is readOnly
	fmt.Println(<-readChannel)
}

func main() {
	ch := make(chan string, 1)
	go ping(ch)
	pong(ch)
	fmt.Println("Done!")
}
