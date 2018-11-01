package main

import (
	"fmt"

	. "github.com/gocity/test_reflect/stu_demo"
)

//"github.com/gocity/test_reflect"

func main() {
	//fmt.Printf("hello, world\n")

	stu := Student{Address: Address{City: "Shanghai", Area: "Pudong"}, Name: "chain", Age: 23}
	StructInfo(stu)

	pos := adder()
	n := 10
	fmt.Println(pos(n))
	fmt.Println(pos(n))

	var j int = 5
	// a 为func 执行返回的 函数
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()
	j *= 2
	a()

}

// 返回一个函数
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
