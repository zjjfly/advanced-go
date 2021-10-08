package ch1

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestString(t *testing.T) {
	//String在运行时实际是reflect.StringHeader
	//这个结构体中的Data指向的底层的byte数组

	s := "Hello"
	//可以把字符串转成StringHeader,然后使用Len来获取字符串长度,但不推荐这种做法
	t.Logf("len(s):%d", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)

	//把字符串转成byte数组可以看它的底层对应的数据,实际是二进制字节序列
	t.Logf("%#v", []byte("Hello,世界"))

	//go字符串的默认编码是UTF-8.
	//如果如果遇到一个错误的UTF8编码输入,将生成一个特别的Unicod字符,之后的错误编码会被忽略
	//错误的编码不会向后扩散,这是UTF-8的良好特性
	t.Log("\xe4\x00\x00\xe7\x95\x8cabc")

	//使用range遍历,这样不会忽略错误编码
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		t.Log(i, c)
	}

	//如果不想进行utf-8编码,可以先把字符串转成byte数组
	for i, c := range "世界abc" {
		t.Log(i, c)
	}

	//go对string和[]rune相互转化提供了特别支持
	//rune实际就是int32,表示unicode编码,只是为了把char和int32进行区分才引入的一个类型别名
	//由于byte和int32不是一种类型,所以这种转换可能会隐含内存重新分配的开销
	t.Logf("%#v\n", []rune("世界"))
	t.Logf("%#v\n", string([]rune{'世', '界'}))
}

//for range迭代字符串的模拟实现
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

func TestForOnString(t *testing.T) {
	forOnString("世界abc", func(i int, r rune) {
		t.Logf("%d:%v", i, r)
	})
}

func str2bytes(s string) []byte {
	bytes := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		//引入一个中间变量是为了保证字符串的只读语义
		c := s[i]
		bytes[i] = c
	}
	return bytes
}

func TestStr2Bytes(t *testing.T) {
	bytes := str2bytes("世界abc")
	assert.Equal(t, []byte{0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x61, 0x62, 0x63}, bytes)
}

func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	//把原byte切片复制到一个新的切片,是为了保证字符串的只读语义
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}

func str2runes(s []byte) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, r)
		s = s[size:]
	}
	return []rune(p)
}

func TestStr2Runes(t *testing.T) {
	runes := str2runes([]byte("世界abc"))
	assert.Equal(t, []rune{19990, 30028, 97, 98, 99}, runes)
}

func runes2str(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

func TestRunes2Str(t *testing.T) {
	str := runes2str(str2runes(str2bytes("世界abc")))
	assert.Equal(t, "世界abc", str)
}
