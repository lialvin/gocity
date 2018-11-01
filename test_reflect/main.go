package main

import (
     . "github.com/gocity/test_reflect/stu_demo"
)

func main() {
	stu := Student{Address: Address{City: "Shanghai", Area: "Pudong"}, Name: "chain", Age: 23}
	StructInfo(stu)
}
