package ch1

import (
	"sync"
	"testing"
)

//有缓冲的channel的第k个接受完成操作发生在第k+c个发送完成之前,c就是channel的缓冲大小
//利用这个特点可以限制最大的goroutine个数
var ch chan int

func work(i int) {
	println(i)
}

func TestLimit(t *testing.T) {
	var wg sync.WaitGroup
	//使用缓冲channel控制同时并发的goroutine数
	ch = make(chan int, 100)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		ch<-1
		go work(i)
		<-ch
		wg.Done()
	}
	wg.Wait()
}
