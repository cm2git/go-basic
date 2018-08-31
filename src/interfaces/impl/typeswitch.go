package impl

import "fmt"

func typeSwitch(i interface{}) {

	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func TypeSwitchDemo() {
	i := 100
	typeSwitch(i)
	s := "hello,go"
	typeSwitch(s)
	f := float64(100)
	typeSwitch(f)
}
