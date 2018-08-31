package impl

import "fmt"

//declare interface
type PersonFinder interface {
	QueryName() string
}

func GetNames(pf []PersonFinder) {
	for _, v := range pf {
		fmt.Println(v.QueryName())
	}
}
