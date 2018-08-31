package main

import "fmt"

func main() {

	iArr := [3]int{1, 2, 3}
	change(iArr[:])
	fmt.Print(iArr)
}

func change(arr []int) {
	arr[0] = 10
}
