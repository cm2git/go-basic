package main

import (
	"fmt"
)

func main() {

	done := make(chan int)
	go send(done)

	for {
		v, ok := <-done
		fmt.Println(v, ok)

		if ok == false {
			break
		}
	}
	fmt.Println("main function")
}

func send(done chan<- int) {
	fmt.Println("prepare send data to the channel done ")
	done <- 1
	done <- 2
	close(done)
	fmt.Println("send data to the channel done ")
}
