package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int64, 1) // make this 2 and see the magic
	go func() {
		fmt.Println("I'm going to write data to channel and now this is going to block this go routines exec until " +
			"data is read from channel")
		c <- 64
		fmt.Println("I have written data to channel once")
		c <- 65
		fmt.Println("I have written data to channel twice")
		c <- 66
		fmt.Println("I have written data to channel thrice")
		//c <- 65
		//c <- 66  // uncomment previous and this line and line 24 and see what happens ;)
		//the next segment need not be executed since now the main go routine is unblocked after the read and if it exists before
		// these prints are exec then they wont be executed obviously!
		fmt.Println("GoodBye!")
	}()
	fmt.Println("Next line is going to block until there is data in channel")
	// unlike before prog doesn't exit even though previous func is a go routine
	runtime.Gosched()
	fmt.Println(<-c)
	time.Sleep(time.Second * 2)
	//fmt.Println(<-c) // event though 66 is not read doesnt matter, there can be n number of writes for unbuffered channel but only n reads,
	//  if n+1 then that routine will block  and will deadlock or break the prog with error if no go routine is running
	fmt.Println("Data read from channel")
	fmt.Println("GoodBye!!!")
}
