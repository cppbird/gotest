package main

import (
	"fmt"
	"net"
	"net/http"
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
	fmt.Println("get one conn", c.RemoteAddr().String())
	fmt.Println("get conn")
	http.NewServeMux()
	for i := 0; i < 10; i++ {
		buf := make([]byte, 10)
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if i == 5 {
			c.Close()
		}
		fmt.Printf("read %s", string(buf))
	}

}
