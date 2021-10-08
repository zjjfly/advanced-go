package ch3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVarAssembly(t *testing.T) {
	assert.Equal(t,9527,Id)
	assert.Equal(t,"gopher",Name)
	assert.Equal(t,"123456",Name1)
	//Name1的问题是如果改了Name1Data,Name1也就随之更改了
	Name1Data[0]='0'
	assert.Equal(t,"023456",Name1)
	assert.Equal(t,"gopher",Name2)
}

func TestFunctionAssembly(t *testing.T) {
	asmPrint()
	t.Log('h','i')
}


