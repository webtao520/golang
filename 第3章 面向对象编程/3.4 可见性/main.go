package main

/*
	Go语言对关键字的增加非常吝啬，其中没有private、protected、public这样的关键
	字。要使某个符号对其他包（package）可见（即可以访问），需要将该符号定义为以大写字母
	开头，如：
*/

type Rect struct {
	X, Y          float64
	Width, Height float64
}

/*
	这样，Rect类型的成员变量就全部被导出了，可以被所有其他引用了Rect所在包的代码访问到。
	成员方法的可访问性遵循同样的规则，例如：
*/
func (r *Rect) area() float64 {
	return r.Width * r.Height
}

/*
	这样，Rect的area()方法只能在该类型所在的包内使用。
	需要注意的一点是，Go语言中符号的可访问性是包一级的而不是类型一级的。在上面的例
	子中，尽管area()是Rect的内部方法，但同一个包中的其他类型也都可以访问到它。这样的可
	访问性控制很粗旷，很特别，但是非常实用。如果Go语言符号的可访问性是类型一级的，少不
	了还要加上friend这样的关键字，以表示两个类是朋友关系，可以访问彼此的私有成员。
*/

func main() {

}
