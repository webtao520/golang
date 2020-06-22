package main

import (
	"fmt"
)

type Integer int

// 不使用go语言面向对象特性 使用面向过程实现
func Integer_Less(a Integer, b Integer) (c bool) { // 面向过程
	//return  a < b
	c = a < b
	return
}

// 在Go语言中，面向对象的神秘面纱被剥得一干二净。对比下面的两段代码：
func (a Integer) Less(b Integer) bool { // 面向对象
	return a < b
}

func main() {
	var a Integer
	a = 3
	if Integer_Less(a, 2) { // 面向过程的用法
		fmt.Println(a, "Integer_Less 2")
	} else {
		fmt.Println("a > b")
	}

	if a.Less(1) { // 面向对象的用法
		fmt.Println(a, "Less 2")
	} else {
		fmt.Println("a > b")
	}
}

/*
“在Go语言中没有隐藏的this指针”这句话的含义是：
 方法施加的目标（也就是“对象”）显式传递，没有被隐藏起来；
 方法施加的目标（也就是“对象”）不需要非得是指针，也不用非得叫this。

*/
