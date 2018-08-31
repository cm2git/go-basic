package main

import "fmt"

func main() {
	i := 10
	pi := &i
	print(i)
	print(pi)

	s := string("hello")
	print(change([]rune(s)))
	print(s)

}

func change(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func print(p ...interface{}) {
	fmt.Println(p)
}
