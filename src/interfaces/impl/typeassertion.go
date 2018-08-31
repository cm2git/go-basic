package impl

import "fmt"

//type assertion
func typeAssertion(i interface{}) {
	value, ok := i.(int)
	fmt.Println(value, ok)
}

func TypeAssertionDemo() {
	s := "hello,go"
	typeAssertion(s)
	i := 100
	typeAssertion(i)
}
