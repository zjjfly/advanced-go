package ch1

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func TestMutexAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	t.Log(total.value)
}

func BenchmarkAtomic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		go worker(&wg)
		go worker(&wg)
		wg.Wait()
	}
}

var total2 uint32

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := uint32(0); i < 100; i++ {
		atomic.AddUint32(&total2, i)
	}
}

func TestBuiltInAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()
	t.Log(total2)
}

func BenchmarkAtomic2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		go worker2(&wg)
		go worker2(&wg)
		wg.Wait()
	}
}

func TestLoadStore(t *testing.T) {
	var config atomic.Value
	loadConfig := func() int {
		return rand.Int()
	}
	config.Store(loadConfig())
	//起一个后台加载最新配置的goroutine
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()
	//起多个处理请求的goroutine
	for i := 0; i < 10; i++ {
		go func(n int) {
			for r := range [...]int{1, 2, 3, 4, 5} {
				c := config.Load()
				t.Log(n, r, c)
			}
		}(i)
	}
}
