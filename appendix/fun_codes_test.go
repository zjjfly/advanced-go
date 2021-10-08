package appendix

import (
	"testing"
)

func TestRandom(t *testing.T) {
	for i := range random(100) {
		t.Log(i)
	}
}
