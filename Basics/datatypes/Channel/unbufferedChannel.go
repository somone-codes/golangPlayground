package main

import "fmt"

func main() {
	c := make(chan int64) // unbuffered / synchronous channel
	//writing a go routine to push data to channel since if done here as channel are blocking , it will block this main
	//go routine , hence pushing it in a func and that go routine will be blocking until the main go routine has completed read
	//and main go routine here blocks at channel read until data is written into the channel
	go func() {
		fmt.Println("I'm going to write data to channel and now this is going to block this go routines exec until " +
			"data is read from channel")
		c <- 64
		//c <- 65
		//c <- 66  // uncomment previous and this line and line 24 and see what happens ;)
		//the next segment need not be executed since now the main go routine is unblocked after the read and if it exists before
		// these prints are exec then they wont be executed obviously!
		fmt.Println("I have written data to channel")
		fmt.Println("GoodBye!")
	}()
	fmt.Println("Next line is going to block until there is data in channel")
	// unlike before prog doesn't exit even though previous func is a go routine
	fmt.Println(<-c)
	//fmt.Println(<-c) // event though 66 is not read doesnt matter, there can be n number of writes for unbuffered channel but only n reads,
	//  if n+1 then that routine will block  and will deadlock or break the prog with error if no go routine is running
	fmt.Println("Data read from channel")
	fmt.Println("GoodBye!!!")
}
