package main 

import (
	"one"
	"two"
)

/*
	这里我们定义了两个接口，一个叫one.ReadWriter，一个叫two.Istream，两者都定义
	了Read()、Write()方法，只是定义次序相反。one.ReadWriter先定义了Read()再定义了
	Write()，而two.IStream反之。
	在Go语言中，这两个接口实际上并无区别，因为：
	 任何实现了one.ReadWriter接口的类，均实现了two.IStream；  任何one.ReadWriter接口对象可赋值给two.IStream，反之亦然；
	 在任何地方使用one.ReadWriter接口与使用two.IStream并无差异。
	以下这些代码可编译通过：
*/
/*
接口赋值并不要求两个接口必须等价。如果接口A的方法列表是接口B的方法列表的子集，
那么接口B可以赋值给接口A。例如，假设我们有Writer接口：
*/
type Writer interface { 
 Write(buf []byte) (n int, err error) 
} 
//就可以将上面的one.ReadWriter和two.IStream接口的实例赋值给Writer接口：
var file1 two.IStream = new(File) 
var file4 Writer = file1 
//但是反过来并不成立：
var file1 Writer = new(File) 
var file5 two.IStream = file1 // 编译不能通过
//这段代码无法编译通过，原因是显然的：file1并没有Read()方法。

var file1 two.IStream = new(File) 
var file2 one.ReadWriter = file1 
var file3 two.IStream = file2

func main(){

}