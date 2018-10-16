package main

import (
	"fmt"
	"strconv"
)

type I interface {
	Foo()
}

type T struct {
	val int64
}

func (t T) Foo() {
	fmt.Println("val: " + strconv.FormatInt(t.val, 10))
	t.val = 1
}

func Foo2(param I) {
	param.Foo()
}

type I2 interface {
	Foox()
}

type T2 struct {
	val int64
}

func (t2 *T2) Foox() {
	fmt.Println("val: " + strconv.FormatInt(t2.val, 10))
	t2.val = 1
}

func Foox2(param *T2) {
	param.Foox()
}

func main() {
	test := T{}
	test.val = 44
	test.Foo()
	Foo2(test)

	test2 := T2{}
	test2.val = 33
	test2.Foox()
	Foox2(&test2)
}
