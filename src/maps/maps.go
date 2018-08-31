package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["aaa"] = 100
	fmt.Println(m)
	mm := map[int]int{
		0: 0,
		1: 1,
	}
	change(mm)
	fmt.Println(mm)
}

func change(m map[int]int) {
	v, ok := m[5]
	fmt.Println(v, ok)
	delete(m, 0)
	v1, ok1 := m[5]
	fmt.Println(v1, ok1)
	//m[0]=100
}
