package ch1

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func Test(t *testing.T) {
	ch := make(chan int)
	go Producer(3, ch)
	go Producer(5, ch)
	go Consumer(ch)

	//time.Sleep(3 * time.Second)

	//捕获系统信号,使得通过ctrl+c退出程序
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,os.Interrupt, os.Kill)
	fmt.Printf("quit (%v)\n", <-sig)
}
