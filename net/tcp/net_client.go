package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	var (
		err  error
		fbs  []byte
		conn net.Conn
	)

	conn, err = net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		return
	}
	defer conn.Close()

	if fbs, err = ioutil.ReadFile("//Users/jdq/go/src/gotest/fuck.mp4"); err != nil {
		log.Printf("ioutil err:%v", err)
	}

	for {
		if _, err := conn.Write(fbs); err != nil {
			fmt.Println("write", err)
			return
		}
		log.Printf("done")
		return
	}
}
