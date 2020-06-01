package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:9941")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		a, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(a)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println("get conn")
	for {
		buf := make([]byte, 10)
		time.Sleep(1 * time.Second)
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
	}
}
