package main

/*
    Go语言的结构体（struct）和其他语言的类（class）有同等的地位，但Go语言放弃了包括继
	承在内的大量面向对象特性，只保留了组合（composition）这个最基础的特性。
	组合甚至不能算面向对象特性，因为在C语言这样的过程式编程语言中，也有结构体，也有
	组合。组合只是形成复合类型的基础。
	上面我们说到，所有的Go语言类型（指针类型除外）都可以有自己的方法。在这个背景下，
	Go语言的结构体只是很普通的复合类型，平淡无奇。例如，我们要定义一个矩形类型：

	type Rect struct {
	x, y float64
	width, height float64
	}

	然后我们定义成员方法Area()来计算矩形的面积：

	func (r *Rect) Area() float64 {
	return r.width * r.height
	}

	可以看出， Go语言中结构体的使用方式与C语言并没有明显不同
*/
func main() {

}
