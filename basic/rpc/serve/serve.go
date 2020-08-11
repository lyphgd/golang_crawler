package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcdemo "v1/basic/rpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})         //rpc包里面的Register方法,传入一个服务结构,客户端可以调用这个结构里面的所有方法
	listener, err := net.Listen("tcp", ":1234") //买电话,插入电话卡
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept() //开机监听，无信号阻塞,有信号过来服务
		if err != nil {
			log.Printf("accept error:%v", err)
			continue
		}
		go jsonrpc.ServeConn(conn) //开一个协程去服务
	}
}
