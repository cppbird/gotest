package main

import (
	"bufio"
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
	reader := bufio.NewReader(c)
	for {
		m, _ := reader.ReadBytes('\n')
		time.Sleep(100 * time.Second)
		fmt.Println(string(m))
	}
}
