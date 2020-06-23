package main 

import (
	"fmt"
)
// 在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口，例如：
type  File struct {
	// ...
}

func (f *File) Read(buf []byte) (n int, err error) 
func (f *File) Write(buf []byte) (n int, err error) 
func (f *File) Seek(off int64, whence int) (pos int64, err error) 
func (f *File) Close() error

/*
  这里我们定义了一个File类，并实现有Read()、Write()、Seek()、Close()等方法。设
  想我们有如下接口：
*/

type IFile interface {
	Read (buf []byte)(n int, err error)
	Write(buf []byte)(n int, err error)
	Seek(off int64, whence int)(pos int64, err error)
	Close() error
}

type IReder interface {
	Read (buf []byte)(n int,err error)
}

type IWriter interface {
	Write(buf []byte)(n int,err error)
}

type ICloser interface {
	Close() error
}

/*
尽管File类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是File类实现了
这些接口，可以进行赋值：
*/
var file1 IFile = new(File) 
var file2 IReader = new(File) 
var file3 IWriter = new(File) 
var file4 ICloser = new(File)


/*
	Go语言的非侵入式接口，看似只是做了很小的文法调整，实则影响深远。
	其一，Go语言的标准库，再也不需要绘制类库的继承树图。你一定见过不少C++、Java、C# 
	类库的继承树图。这里给个Java继承树图：
	http://docs.oracle.com/javase/1.4.2/docs/api/overview-tree.html 
	在Go中，类的继承树并无意义，你只需要知道这个类实现了哪些方法，每个方法是啥含义
	就足够了。
	其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才
	合理。接口由使用方按需定义，而不用事前规划。
	其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦
	合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口。
*/


func main (){
   
}