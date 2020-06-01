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

	for {
		if _, err := conn.Write([]byte("foo")); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
		return
	}
	time.Sleep(5000 * time.Second)
	fmt.Println("ok")

}
