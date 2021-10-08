package ch5

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	msg := "Hello,World"
	resp, err := http.Post("http://localhost:8080/", "application/text", strings.NewReader(msg))
	if err != nil {
		t.Fail()
	}
	defer resp.Body.Close()
	respMsg, err := ioutil.ReadAll(resp.Body)
	assert.Equal(t, msg,string(respMsg))
}

