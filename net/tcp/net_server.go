package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
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
		reader := bufio.NewReader(c)
		bytes, _ := reader.Read(p)
		ioutil.WriteFile("fuck.mp4", bytes, 0644)
		log.Printf("fuck.mp4 done")
		return
	}
}
