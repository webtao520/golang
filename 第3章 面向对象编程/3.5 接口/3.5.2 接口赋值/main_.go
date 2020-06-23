package main 

import (
	"fmt"
)
/*
	接口赋值在Go语言中分为如下两种情况：
	 将对象实例赋值给接口；
	 将一个接口赋值给另一个接口。
	先讨论将某种类型的对象实例赋值给接口，这要求该对象实例实现了接口要求的所有方法，
	例如之前我们作过一个Integer类型，如下：
*/

type  Integer int 

func (a Integer)Less(b Integer)bool {
	return a < b
}

func (a *Integer)Add(b Integer){
	*a+=b
}
// 相应地，我们定义接口LessAdder，如下：
type LessAdder  interface {
	Less(b Integer)bool
	Add(b Integer)
}

/*
现在有个问题：假设我们定义一个Integer类型的对象实例，怎么将其赋值给LessAdder
接口呢？应该用下面的语句(1)，还是语句(2)呢？
*/
var a Integer =1
var b LessAdder=&a // ...(1)
var b LessAdder=a // ...(2)
// 答案是应该用语句(1)。原因在于，Go语言可以根据下面的函数：
func (a Integer)Less(b Integer)bool
// 自动生成一个新的Less()方法
func (a *Integer)Less(b Integer)bool {
	return (*a).Less(b)
}
/*
这样，类型*Integer就既存在Less()方法，也存在Add()方法，满足LessAdder接口。而
从另一方面来说，根据
*/
func (a *Integer) Add(b Integer) 
//这个函数无法自动生成以下这个成员方法：
func (a Integer) Add(b Integer) { 
 (&a).Add(b)
}
/*
因为(&a).Add()改变的只是函数参数a，对外部实际要操作的对象并无影响，这不符合用
户的预期。所以，Go语言不会自动为其生成该函数。因此，类型Integer只存在Less()方法，
缺少Add()方法，不满足LessAdder接口，故此上面的语句(2)不能赋值。
为了进一步证明以上的推理，我们不妨再定义一个Lesser接口，如下：
*/

type Lesser interface { 
	Less(b Integer) bool
   } 
   //然后定义一个Integer类型的对象实例，将其赋值给Lesser接口：
   var a Integer = 1 
   var b1 Lesser = &a //... (1) 
   var b2 Lesser = a //... (2) 
   /*
   正如我们所料的那样，语句(1)和语句(2)均可以编译通过。
   我们再来讨论另一种情形：将一个接口赋值给另一个接口。在Go语言中，只要两个接口拥
   有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值。
   下面我们来看一个示例，这是第一个接口：
 */  

func main (){

}