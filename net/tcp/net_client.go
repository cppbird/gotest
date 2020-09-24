package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var (
		err  error
		conn net.Conn
	)

	conn, err = net.Dial("tcp", "127.0.0.1:9941")
	if err != nil {
		return
	}
	fmt.Println("dial ok", conn.RemoteAddr().String())
	defer func() {
		conn.Close()
		fmt.Println("close conn")
	}()

	for i := 0; i < 10; i++ {
		if _, err := conn.Write([]byte("foo")); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("write foo")
		time.Sleep(1 * time.Second)
	}
}
