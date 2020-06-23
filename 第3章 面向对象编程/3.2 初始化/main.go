package main

import (
	"fmt"
)

type Rect struct {
	x,y float64
	width,height float64
}

func (r *Rect)Area()float64{
	return r.width * r.height
}

/*
在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以
NewXXX来命名，表示“构造函数”：
*/
func NewRect(x,y,width,height float64) *Rect{
    return &Rect{x,y,width,height}
}
/*
    这一切非常自然，开发者也不需要分析在使用了new之后到底背后发生了多少事情。在Go
	语言中，一切要发生的事情都直接可以看到。
	在Go语言中，未进行显式初始化的变量都会被初始化为该类型的零值，例如bool类型的零
	值为false，int类型的零值为 0，string类型的零值为空字符串。
*/
func main(){
   rect1:=new (Rect)
   fmt.Println(rect1)
   rect2:=&Rect{0,0,100,200}
   rect3:=&Rect{width: 100, height: 200}
   fmt.Println(rect1,rect2,rect3)
   fmt.Println(NewRect(0,0,0,0))
}