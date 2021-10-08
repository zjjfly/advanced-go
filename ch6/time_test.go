package ch6

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	i := time.Now().UnixNano() / 1e6
	t.Log(i)
}

