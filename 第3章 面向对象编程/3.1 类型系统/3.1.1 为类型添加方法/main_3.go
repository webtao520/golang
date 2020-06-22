package main

import (
	"fmt"
)

/*
	Go语言中的面向对象最为直观，也无需支付额外的成本。如果要求对象必须以指针传递，
	这有时会是个额外成本，因为对象有时很小（比如4字节），用指针传递并不划算。
	只有在你需要修改对象的时候，才必须用指针。它不是Go语言的约束，而是一种自然约束。
	举个例子：

*/

type Integer int

// 引用传递
func (a *Integer) Add(b Integer) {
	*a += b
}

// 值传递
func (a Integer) change(b Integer) {
	a += b
}

func main() {
	var a Integer = 1
	//a.Add(2)
	a.change(3)
	//fmt.Println("引用传递:\n a=",a)
	fmt.Println("值传递:\n", a)
}
