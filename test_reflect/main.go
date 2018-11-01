package main

import (
	"github.com/gocity/test_reflect"
)

func main() {
	stu := Student{Address: Address{City: "Shanghai", Area: "Pudong"}, Name: "chain", Age: 23}
	StructInfo(stu)
}
