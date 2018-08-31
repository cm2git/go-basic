package main

import (
	"fmt"
	"sync"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	processDemo()
}

func processDemo() {
	no := 3
	var wg sync.WaitGroup
	fmt.Println("wg zero value is ", wg)
	for i := 0; i < no; i++ {
		wg.Add(i)
		go process(i, &wg)
	}
	wg.Wait()
}

func writeDemo() {
	ch := make(chan int, 2)
	go write(ch)
	//time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		//time.Sleep(2 * time.Second)

	}
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
