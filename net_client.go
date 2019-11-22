package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		if _, err := conn.Write([]byte("send")); err != nil {
			fmt.Println("write", err)
			return
		}
		// buf := make([]byte, 1000)
		time.Sleep(1 * time.Second)
		// if _, err = conn.Read(buf); err != nil {
		// 	fmt.Println("read", err)
		// 	return
		// }

	}

}
