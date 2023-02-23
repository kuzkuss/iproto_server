package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "localhost:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer func(openConn *net.TCPConn) {
		err := openConn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(conn)
}
