package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:8000")
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

	for {
		b := make([]byte, 100)
		fmt.Println(c.Read(b))
		c.Write([]byte("test"))
	}
}
