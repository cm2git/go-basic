package main

import (
	"fmt"
)

func main() {
	//done := make(chan string)
	//go hello(done)
	//data := <-done
	//fmt.Println(data)
	//time.Sleep(1*time.Second)
	//fmt.Println("main goroutine")
	demo()
	bufferedChannel()
}

func hello(done chan<- string) {
	fmt.Println("hello,Goroutine")
	done <- "hello,Go"
}

func demo() {
	done := make(chan int)
	go sender(done)
	receiver(done)
}

func sender(done chan<- int) {
	for i := 1; i < 10; i++ {
		done <- i
	}
	close(done)
}

func receiver(done chan int) {

	/*for v := range done {
		fmt.Println(v)
	}*/

	for {

		v, ok := <-done
		if ok == false {
			fmt.Println("=====================")
			break
		}
		fmt.Println(v, ok)
	}
}

func bufferedChannel() {

	ints := make(chan int, 3)
	ints <- 1
	ints <- 2
	ints <- 3

	for i := 0; i < 3; i++ {
		fmt.Println(<-ints)
	}
}
