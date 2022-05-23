package main

import (
	"fmt"
)

func main() {

	// DoLoopTest()
	// DoArrPointerStructTest()

	config := &ArrayMe{
		IntMe: 1,
	}

	config.SetInt()
	fmt.Println("config", config.IntMe)

}

func (a *ArrayMe) SetInt() {
	b := *a
	fmt.Println("B", b.IntMe)
	b.IntMe = 5
	fmt.Println("B", b.IntMe)
}

func DoLoopTest() {
	var list []int

	// list = &[]int{1, 2, 3, 4, 5}

	for _, v := range list {
		fmt.Println(v)
	}
}

type ArrayMe struct {
	IntMe int
}

func DoArrPointerStructTest() {
	arrayMe := []*ArrayMe{}

	for _, v := range arrayMe {
		fmt.Println(v)
	}
}
