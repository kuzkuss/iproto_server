package main

import (
	"log"
	"net"

	"github.com/kuzkuss/iproto_server/internal/delivery"
	"github.com/kuzkuss/iproto_server/internal/repository"
	"github.com/kuzkuss/iproto_server/internal/usecase"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	rep := repository.New()
	useCase := usecase.New(rep)
	del := delivery.New(useCase)

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(del, conn)
	}
}

func handleConnection(del delivery.DeliveryI, conn *net.TCPConn) {
	defer conn.Close()
	req, err := del.GetRequest(conn)
	if err != nil {
		log.Println(err)
		return
	}

	err = del.HandleRequest(req)
	if err != nil {
		log.Println(err)
		return
	}

	if err := del.SendResponse(conn); err != nil {
		log.Println(err)
		return
	}
}

