package one

/*
import (
	"fmt"
)
*/
type ReadWrite interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

/*
func main (){

}
*/
