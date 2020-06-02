package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println("与客户端建立连接 :", addr)
	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("客户端连接异常:", err)
			break
		}
		msg := string(bytes[:n])
		if msg == "exit" {
			break
		}
		fmt.Println(addr, ":", msg)
	}
	fmt.Println(addr, "断开连接")
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("服务端启动完成.....")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error message :", err)
		}
		go process(conn)
	}
}
