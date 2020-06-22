package main 

import (
	"fmt"
)

type Integer int

// 不使用go语言面向对象特性 使用面向过程实现
func Integer_Less (a  Integer, b Integer) (c bool) {
	//return  a < b
	c=a <b
	return
}

func main (){
	 var a Integer 
	 a =3
	 if Integer_Less(a,2) {
		 fmt.Println(a, "Integer_Less 2")
	 }else {
		 fmt.Println("a > b")
	 }
}