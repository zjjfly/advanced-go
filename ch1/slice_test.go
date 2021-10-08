package ch1

/*
	slice比数组更常用,因为它的长度是动态的,而数组的长度是固定的,且长度是数组类型的一部分
	slice还有一个优势是在参数传递的时候不会复制底层的数据,而数组会
*/

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"sort"
	"testing"
	"unsafe"
)

func TestSliceAppend(t *testing.T) {
	//使用内置函数append在切片后追加元素
	var a []int
	a = append(a, 1)
	//追加切片,需要解包
	a = append(a, []int{2, 3, 4}...)
	assert.Equal(t, []int{1, 2, 3, 4}, a)
	//append也可以在切片头追加元素,但这样可能会引起内存重新分配,影响性能
	a = append([]int{0}, a...)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, a)

	//利用append支持链式操作的特点,在slice中间插入元素或slice
	a = append(a[:2], append([]int{-1}, a[2:]...)...)
	assert.Equal(t, []int{0, 1, -1, 2, 3, 4}, a)
	a = append(a[:3], append([]int{5, 6}, a[3:]...)...)
	assert.Equal(t, []int{0, 1, -1, 5, 6, 2, 3, 4}, a)
	//这种方法有一个弊端是会产生临时slice,使用copy可以避免
	//copy实现插入元素
	a = append(a, 0)
	copy(a[3:], a[2:])
	a[2] = -2
	assert.Equal(t, []int{0, 1, -2, -1, 5, 6, 2, 3, 4}, a)
	//也可以实现插入切片
	x := []int{-4, -3}
	a = append(a, x...)       // 为x切片扩展足够的空间
	copy(a[2+len(x):], a[2:]) // a[i:]向后移动len(X)个位置
	copy(a[2:], x)
	assert.Equal(t, []int{0, 1, -4, -3, -2, -1, 5, 6, 2, 3, 4}, a)
}

func TestDelete(t *testing.T) {
	//删除slice尾部元素
	a := []int{1, 2, 3}
	a = a[:len(a)-1]
	assert.Equal(t, []int{1, 2}, a)

	//删除slice开头元素
	a = a[1:]
	assert.Equal(t, []int{2}, a)
	//使用append完成删除开头元素
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...)
	assert.Equal(t, []int{2, 3}, a)
	//使用copy完成删除开头元素
	a = []int{1, 2, 3}
	a = a[:copy(a, a[1:])]
	assert.Equal(t, []int{2, 3}, a)

	//删除中间元素,使用append或copy实现
	a = []int{1, 2, 3, 4, 5}
	a = append(a[:2], a[3:]...)
	assert.Equal(t, []int{1, 2, 4, 5}, a)
	a = a[:2+copy(a[2:], a[3:])]
	assert.Equal(t, []int{1, 2, 5}, a)
}

//实现slice高效操作的关键是降低内存分配次数,也即是保证append操作不会让slice的Len超过Cap
func TestSliceOptimize(t *testing.T) {
	//slice操作高效的关键是减少内存分配的次数
	//实现一个TrimSpace,使用长度为0但capacity不为0的slice,这样会非常高效
	s := []byte(" ass ")
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	assert.Equal(t, "ass", string(b))

	//删除元素的时候,最好在之前把要删除的元素设为nil
	a := []*os.File{os.Stdout, os.Stderr}
	a[1] = nil
	a = a[:1]
	assert.Equal(t, []*os.File{os.Stdout}, a)
}

//slice的底层byte数组是放在内存中的,当一段很长的slice中的一小部分被引用时,整个slice都不会被释放
func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	//解决办法是把这一小段内存追加到一个空slice中
	return append([]byte{}, b...)
}

//slice的实际类型是reflect.SliceHeader
func TestForceConvert(t *testing.T) {
	var a = []float64{4.1, 2.3, 5.4, 7.2, 2.7, 1.9, 88, 1.8}
	//不同类型的slice之间一般不会相互转换,除了某些情况

	//比如对float切片排序,为了排序速度,可以把他先转换成int切片
	var b = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(
		a):cap(a)]
	sort.Ints(b)
	assert.Equal(t, []float64{1.8, 1.9, 2.3, 2.7, 4.1, 5.4, 7.2, 88}, a)
	//第二种写法
	a = []float64{4.1, 2.3, 5.4, 7.2, 2.7, 1.9, 88, 1.8}
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr
	sort.Ints(c)
	assert.Equal(t, []float64{1.8, 1.9, 2.3, 2.7, 4.1, 5.4, 7.2, 88}, a)
}

var a []float64

func init() {
	for i := 0; i < 2000000; i++ {
		a = append(a, rand.Float64()+rand.Float64()*100.0)
	}
}

func BenchmarkSortFloatSlice1(b *testing.B) {
	b.ResetTimer()
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr
	sort.Ints(c)
}

func BenchmarkSortFloatSlice2(b *testing.B) {
	b.ResetTimer()
	sort.Float64s(a)
}
