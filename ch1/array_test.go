package ch1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var a [3]int
	assert.Equal(t, "[0 0 0]", fmt.Sprint(a))
	var b = [...]int{1, 2, 3}
	assert.Equal(t, "[1 2 3]", fmt.Sprint(b))
	var c = [...]int{2: 3, 1: 2}
	assert.Equal(t, "[0 2 3]", fmt.Sprint(c))
	var d = [...]int{1, 2, 4: 5, 6}
	assert.Equal(t, "[1 2 0 0 5 6]", fmt.Sprint(d))
}

func TestArrayPointer(t *testing.T) {
	//数组会在传递的时候复制,所以会带来的开销
	//避免这种开销的方法是使用数组指针
	var a = [...]int{1, 2, 3}
	var b = &a
	//操作数组指针和操作数组的方式基本一致
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])
	for i, v := range b {
		fmt.Println(i, v)
	}
	//不同长度的数组得到的指针的类型是不同的
}

func TestZeroLenArray(t *testing.T) {
	//空数组占用的内存大小是0
	//如果只是需要把channel作为同步工具而不关系channel中传输的类型,可以使用空数组,减少内存开销
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	//实际开发中更倾向于使用空struct
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<-c2
}
