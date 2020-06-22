package main 

import (
	"fmt"
)

//  在Go语言中，你可以给任意类型（包括内置类型，但不包括指针类型）添加相应的方法，
type  Integer int 

func  (a Integer)Less(b Integer)bool{
	//fmt.Println(a)
	return a<b
}

/*
		在这个例子中，我们定义了一个新类型Integer，它和int没有本质不同，只是它为内置的
		int类型增加了个新方法Less()。
		这样实现了Integer后，就可以让整型像一个普通的类一样使用：
*/
func main () {
   var a Integer
   a=1
   if a.Less(2){
	fmt.Println(a,"Less 2") 
   }else {
	 fmt.Println("a>b")
   }
}