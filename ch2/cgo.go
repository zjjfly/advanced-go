package ch2

/*
#include <stdio.h>
static const char* cs = "hello";

void printint(int v) {
	printf("printint: %d\n", v);
}
*/
import "C"
import (
	"github.com/zjjfly/advanced-go/ch2/cgo_helper"
)

//上面导入的这个包C实际是一个虚包，实际cgo在这个包中找到的所有c函数的字段
//不同的go package中的C包中的类型是不能通用的
func PrintInt() {
	i := 1
	C.printint(C.int(i))
}

func PrintCString() {
	str := "Hello,World"
	//s的类型是当前C包中的char类型
	//所以下面这行代码会在编译的时候出错
	//s := C.CString(str)
	//cgo_helper.PrintCString(s)
	cgo_helper.PrintCString2(str)
}
