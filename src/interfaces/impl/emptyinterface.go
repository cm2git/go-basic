package impl

import "fmt"

func emptyInterface(i interface{}) {
	fmt.Printf("type of i is %T ,value of i is %v\n", i, i)
}

func emptyInterfaces(ii ...interface{}) {
	fmt.Printf("type of ii is %T\n", ii) //[]interface{}
	for _, v := range ii {
		fmt.Printf("type of i is %T ,value of i is %v\n", v, v)
		v = 100
	}
	fmt.Println(len(ii))
}

func EmptyInterfaceDemo() {
	i := 10
	emptyInterface(i)
	emptyInterfaces(i, "hello", "GO")
	emptyInterfaces(i, 2, 3)
	fmt.Print(i)

	ii := []string{"aa", "bb"}
	emptyInterfaces(ii)

	num := []int{89, 90, 95}
	add(num...)
}

func add(num ...int) {
	fmt.Printf("the type of num is %T,value is %v", num, num)
}
