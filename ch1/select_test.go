package ch1

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

//使用select实现管道的超时判断
func TestTimeout(t *testing.T) {
	in := make(chan string)
	select {
	case v := <-in:
		fmt.Println(v)
	case <-time.After(time.Second):
		return
	}
}

//使用select实现非阻塞的管道发送或接收
func TestNoneBlockingSend(t *testing.T) {
	in := make(chan string)
	select {
	case v := <-in:
		fmt.Println(v)
	default:
		fmt.Println("HeHe")
	}
}

//当有多个通道可操作时,select会随机选择一个通道,基于这个特性可以实现随机数生成
func TestRandom(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:

			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func TestTerminateGoroutine(t *testing.T) {
	var work = func(ch chan bool) {
		for {
			select {
			default:
				println("Working...")
			case <-ch:
				return
			}
		}
	}
	ch := make(chan bool)
	go work(ch)
	time.Sleep(time.Second)
	ch <- true
}

func TestSafeTerminate(t *testing.T) {
	var work = func(wg *sync.WaitGroup, cannel chan bool) {
		defer wg.Done()
		for {
			select {
			default:
				println("Working...")
			case <-cannel:
				return
			}
		}
	}
	cannel := make(chan bool)
	//WaitGroup用于保证退出goroutine的清理工作完成之后再终止main线程
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(&wg, cannel)
	}
	time.Sleep(time.Second)
	//close会向这个被关闭的通道的所有接收操作发送一个0h或一个可选的失败标志
	close(cannel)
}

//使用context包提供的工具来实现安全退出
func TestContext(t *testing.T) {
	var work = func(cxt context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-time.After(1 * time.Second):
				println("Working...")
			case _,ok :=<-cxt.Done():
				//ctx.Done返回的channel会在调用cancel之后关闭,而关闭的channel是可以读取数据的,并且读取不会被阻塞
				t.Log(ok)
				return
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(ctx, &wg)
	}
	time.Sleep(3*time.Second)
	cancel()
	wg.Wait()
}

func TestSa(t *testing.T) {
	var worker = func() <-chan int{
		ints := make(chan int)
		go func() {
			time.Sleep(1*time.Second)
			close(ints)
		}()
		return ints
	}
	select {
	case <-worker():
		t.Log("heh")
		case <-time.After(3*time.Second):
			return
	}
}

