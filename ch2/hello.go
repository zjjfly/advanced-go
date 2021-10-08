//+build go1.10

package ch2

/*
#include <stdio.h>

static void SayHello1(const char* s) {
	puts(s);
}

void SayHello2(const char* s);
void SayHello3(const char* s);
void SayHello4(char* s);
//cgo内置了_GoString_这个c类型,用于支持go语言中的字符串
void SayHello5(_GoString_ s);
*/
import "C"
import "fmt"

//为了和cgo导出的函数适配,函数的参数不能加const

//export SayHello4
func SayHello4(s *C.char) {
	fmt.Print(C.GoString(s))
}

//export SayHello5
func SayHello5(s string) {
	fmt.Print(s)
}

func SayHello1() {
	C.SayHello1(C.CString("Hello, World 1\n"))
}

func SayHello2() {
	C.SayHello2(C.CString("Hello, World 2\n"))
}

func SayHello3() {
	C.SayHello3(C.CString("Hello, World 3\n"))
}

func SayHelloFour() {
	C.SayHello4(C.CString("Hello, World 4\n"))
}

func SayHelloFive() {
	C.SayHello5("Hello, World 5\n")
}
