package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

var quitSemaphore = make(chan bool, 1)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("dial err :%v", err)
		return
	}
	defer conn.Close()

	fmt.Println("connected!")
	b, _ := ioutil.ReadFile("/Users/jdq/Desktop/WechatIMG268.jpeg")
	fmt.Println(len(b))
	fmt.Println(conn.Write(b))
	return
	// go onMessageRecived(conn)

	// // 控制台聊天功能加入
	// for {
	// 	select {
	// 	case <-quitSemaphore:
	// 		fmt.Println("tcp client quit")
	// 		return
	// 	default:
	// 		var msg string
	// 		fmt.Scanln(&msg)
	// 		if msg == "quit" {
	// 			break
	// 		}
	// 		b := []byte(msg + "\n")
	// 		conn.Write(b)
	// 	}
	// }
}

func onMessageRecived(conn *net.TCPConn) {
	// reader := bufio.NewReader(conn)
	var msg = make([]byte, 1000)
	for {
		// msg, err := reader.ReadString('\n')
		_, err := conn.Read(msg)
		fmt.Println(string(msg))
		if err != nil {
			fmt.Println("io reader ,err :%v", err)
			quitSemaphore <- true
			break
		}
	}
}
