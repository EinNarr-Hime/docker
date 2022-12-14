package main

import (
	"fmt"
	"time"
)

func say(c chan bool) {
	for i := 0; i < 5; i++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(i)
	}
	c <- true
}

func main() {
	channel := make(chan bool)
	go say(channel)

	b := <- channel
	fmt.Println(b)
}