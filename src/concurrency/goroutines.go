package main

import (
	"fmt"
	"time"
)

func main() {

	go numbers()
	go alphabets()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main function")

}

func numbers() {

	for i := 1; i <= 5; i++ {
		time.Sleep(101 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}
