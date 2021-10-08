package cgo_helper

//#include <stdio.h>
import "C"

type CChar C.char

func (p *CChar) GoString() string {
	return C.GoString((*C.char)(p))
}

//这个函数设计的不好,因为每个包中导入的C包都是不一样的,这个函数只能够供相同的包中的代码调用
func PrintCString(cs *C.char) {
	C.puts(cs)
}

//这个函数就比上面的那个更好
func PrintCString2(s string)  {
	C.puts(C.CString(s))
}
