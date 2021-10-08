package ch4

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/rpc"
	"strings"
	"testing"
)

func TestRpcClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fail()
	}
	var reply string
	err = client.Call(HelloServiceName+".Hello", "jjzi", &reply)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "hello:jjzi", reply)
}

func TestWrappedRpcClient(t *testing.T) {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		t.Fail()
	}
	var reply string
	err = client.Hello("jjzi", &reply)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "hello:jjzi", reply)
}

func TestJsonRpc(t *testing.T) {
	client, err := JsonDialHelloService("tcp", "localhost:12345")
	if err != nil {
		t.Fail()
	}
	var reply string
	err = client.Hello("jjzi", &reply)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "hello:jjzi", reply)
}

func TestHttpJsonRpc(t *testing.T) {
	var clientResponse struct {
		Id     uint64           `json:"id"`
		Result string `json:"result"`
		Error  interface{}      `json:"error"`
	}
	resp, err := http.Post("http://localhost:8080/json-rpc", "application/json", strings.NewReader("{\"method\":\"HelloService.Hello\",\"params\":[\"jjzi\"],\"id\":0}"))
	defer resp.Body.Close()
	if err != nil {
		t.Fail()
	}
	err = json.NewDecoder(resp.Body).Decode(&clientResponse)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t,"hello:jjzi",clientResponse.Result)
}

func TestProtoRpc(t *testing.T) {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		t.Fail()
	}
	var reply String
	err = client.ProtoHello(&String{Value: "jjzi"}, &reply)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "hello:jjzi", reply.Value)
}

