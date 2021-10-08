package ch1

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"syscall"
	"testing"
)

func TestErrno(t *testing.T) {
	//系统调用返回的错误的类型是syscall.Errno
	err := syscall.Chmod("invalid path", 0666)
	assert.Error(t, err)
	assert.Equal(t,"no such file or directory",err.Error())
}

//处理panic的标准方法,在defer的函数中调用recover
func TestPanicRecover(t *testing.T) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("Unknown panic (%v)\n", r)
			}
		}
		t.Log(err)
	}()
	panic("TODO")
}
