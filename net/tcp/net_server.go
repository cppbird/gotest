package main

import (
	"fmt"
	"net"
	"os"
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
	defer c.Close()

	for {
		tmp := make([]byte, 10)
		_, err := c.Read(tmp)
		if err != nil {
			fmt.Printf("err:%v", err)
			os.Exit(1)
		}
		fmt.Println(5)
	}

}
