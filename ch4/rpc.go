package ch4

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
	定义一个rpc一般分为三部分:服务名,服务要实现的方法列表,注册该类型服务的函数
*/

const HelloServiceName = "HelloService"
const ProtoHelloServiceName = "ProtoHelloService"

func init() {
	RegisterHelloService(new(HelloService))
	RegisterProtoHelloService(new(ProtoHelloService))
	go PlainRpcServer()
	go JsonRpcServer()
	go HttpJsonRpcServer()
}

type HelloServiceInterface interface {
	//rpc服务的方法必须是有两个可序列化的参数,第二个参数必须是指针,方法的返回值必须是error,方法必须是公开的
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	//注册服务,传入的svc中的所有的符合规范的方法都会注册到对应的HelloServiceName服务空间下
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (helloService *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

//定义客户端
type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

//json编码客户端
func JsonDialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: client}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func (p *HelloServiceClient) ProtoHello(request *String, reply *String) error {
	return p.Client.Call(ProtoHelloServiceName+".Hello", request, reply)
}

func PlainRpcServer() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}

func JsonRpcServer() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func HttpJsonRpcServer() {
	http.HandleFunc("/json-rpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     writer,
			ReadCloser: request.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":8080", nil)
}

//使用protobuf生成的类型作为rpc服务的参数和返回
type ProtoHelloServiceInterface interface {
	Hello(request *String, reply *String) error
}

type ProtoHelloService struct {
}

func (service *ProtoHelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func RegisterProtoHelloService(svc ProtoHelloServiceInterface) error {
	//注册服务,传入的svc中的所有的符合规范的方法都会注册到对应的HelloServiceName服务空间下
	return rpc.RegisterName(ProtoHelloServiceName, svc)
}
