package ch1

import (
	"sync"
	"testing"
)

//使用同步原语保证在不同goroutine的操作的顺序性

//使用channel进行同步
func TestChannel(t *testing.T) {
	done := make(chan struct{})
	go func() {
		t.Log("Hello,World")
		done <- struct{}{}
	}()
	<-done
}

//属于Mutex同步
func TestMutex(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		t.Log("Hello,World")
		mu.Unlock()
	}()
	mu.Lock()
}
