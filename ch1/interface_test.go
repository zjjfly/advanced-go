package ch1

/*
	go中,只要一个类型实现了接口包含的方法(不需要实现接口的所有方法),就算是这个接口的实现,类似duck type
*/

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

//使用嵌入接口来伪造私有的private方法
type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func TestInterfaceImpl(t *testing.T) {
	//TB只继承了testing.TB的规范,具体的实现在运行时注入
	var tb testing.TB = new(TB)
	tb.Fatal("Hello") //TB.Fatal disabled!
}

//go对接口类型的转换非常灵活
func TestImplicitConvert(t *testing.T) {
	f, err := os.Open("../glide.yaml")
	if err != nil {
		t.Fail()
	}
	var a io.ReadCloser = f
	var b io.Reader = a
	var c io.Closer = a
	var d = c.(io.Reader)
	var bytes = make([]byte, 1024)
	read, _ := b.Read(bytes)
	str1 := bytes2str(bytes[:read])
	bytes =  make([]byte, 1024)
	f.Seek(0,0)
	read, _ = d.Read(bytes)
	str2 := bytes2str(bytes[:read])
	assert.Equal(t, str1, str2)
}
