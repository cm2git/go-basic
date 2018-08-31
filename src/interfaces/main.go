package main

import (
	"fmt"
	"interfaces/impl"
)

func main() {
	//personDemo()
	//impl.TypeAssertionDemo()
	//impl.TypeSwitchDemo()
	impl.EmptyInterfaceDemo()
}

func personDemo() {
	m := impl.Man{
		Name: "jack",
	}
	fmt.Println(m.QueryName())
	pf := []impl.PersonFinder{m, &impl.Woman{}}
	impl.GetNames(pf)
}
