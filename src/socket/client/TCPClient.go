package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:8080")
	if e != nil {
		panic(e)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		str, e := reader.ReadString('\n')
		msg := str[:len(str)-1]
		if e != nil {
			fmt.Println("error message :", e)
			break
		}
		fmt.Fprint(conn, msg)
		if msg == "exit" {
			break
		}
	}
}
