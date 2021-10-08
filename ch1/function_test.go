package ch1

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"os"
	"testing"
)

func TestFunc(t *testing.T) {
	//函数分为具名函数和匿名函数
	//一个匿名函数的例子
	var Add = func(x int, y int) int {
		return x + y
	}
	assert.Equal(t, 3, Add(1, 2))

	//函数的参数可以是可变的,但需要作为最后一个参数
	var Sum = func(a int, more ...int) int {
		//可变参数实际是一个slice,所以可以使用range进行遍历
		for _, i := range more {
			a += i
		}
		return a
	}
	assert.Equal(t, 10, Sum(1, 2, 3, 4))

	//返回值也可以命名,这样可以直接通过这个名字来修改范围值
	var Inc = func() (v int) {
		defer func() { v++ }()
		return 42
	}
	assert.Equal(t, 43, Inc())

	//使用外部变量的匿名函数叫做闭包
	//闭包是通过引用去访问外部变量的,所以有可能会有问题,这在js中已经被讨论了很多
	for i := 0; i < 3; i++ {
		//在for循环中使用defer是不推荐的做法
		defer func() { t.Log(i) }()
	}
	//输出:
	//3
	//3
	//3
}

func TestSideEffect(t *testing.T) {
	//go函数的参数只复制其中固定的部分,如果其中有指针,那么指针对应的内容是不会复制的,所以可以在函数中修改
	//所以下面的函数可以修改原slice的内容
	f := func(x []int) {
		for i := range x {
			x[i] *= 2
		}
	}

	nums := []int{1, 2, 3}
	f(nums)
	assert.Equal(t, []int{2, 4, 6}, nums)
}

func TestMethod(t *testing.T) {
	//类型的函数就叫方法
	//方法可以通过方法表达式来变成普通函数
	var CloseFile = (*os.File).Close
	f, e := os.Open("slice_test.go")
	if e != nil {
		t.Error("file not found")
	}
	defer CloseFile(f)

	//但方法表达式得到的方法和具体类型是绑定的,无法和其他有相同方法的类型无缝适配
	//一种解决方法是使用方法值
	var ReadAt = f.ReadAt
	var Close = f.Close
	var b []byte
	_, err := ReadAt(b, 0)
	Close()
	if err != nil {
		t.Error(err)
	}
}

func TestInheritance(t *testing.T) {
	type Point struct {
		X, Y float64
	}
	//一个继承Point的类
	type ColoredPoint struct {
		Point
		Color color.RGBA
	}

	var cp ColoredPoint
	cp.X = 1
	t.Log(cp.Point.X) // "1"
	cp.Point.Y = 2
	t.Log(cp.Y) // "2"

}
