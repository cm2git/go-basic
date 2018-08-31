package main

import "fmt"

type Person struct {
	name string
	age  int
	animal
}
type animal struct {
	name string
}

var Animal struct {
	name string
}

func main() {

	//	basic()
	p := Person{
		name: "jack",
	}
	fmt.Println(p)
	change(&p)
	fmt.Println(p)

}

func change(p *Person) {
	p.name = "change Person.name" //<=>(*p).name
}

func basic() {
	p := Person{
		name:   "jack",
		age:    20,
		animal: animal{name: "cat"},
	}
	//p.animal=animal{name:"cat"}
	fmt.Println(p)

	p1 := Person{"rose", 20, animal{name: "cat"}}
	fmt.Println(p1)

	Animal.name = "cat"

	fmt.Println(&Animal)

	var anonymous = struct {
		string
	}{
		string: "anonymous",
	}

	fmt.Println(&anonymous)
}
