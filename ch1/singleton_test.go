package ch1

//在性能敏感的情况下,增加一个数字型的标志位,并使用atomic检查它的位状态
//这比使用互斥锁的性能高很多

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Singleton() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

func TestSingleton(t *testing.T) {
	s := Singleton()
	assert.NotNil(t, s)
}

//把上面的通用代码提取出来,就是go中的sync.Once的实现
type MyOnce struct {
	sync.Mutex
	done uint32
}

func (o *MyOnce) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.Lock()
	defer o.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func TestMyOnce(t *testing.T) {
	o := MyOnce{}
	var n int
	for i := 0; i < 100; i++ {
		go o.Do(func() {
			n += 1
		})
	}
	assert.Equal(t, 1, n)
}

var once sync.Once

//使用标准库的sync.Once实现单例模式
func Singleton2() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func TestSingleton2(t *testing.T) {
	s := Singleton2()
	assert.NotNil(t, s)
}
