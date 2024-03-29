package main

import (
	"log"
	"net"

	"github.com/kuzkuss/iproto_server/internal/delivery"
	"github.com/kuzkuss/iproto_server/internal/repository"
	"github.com/kuzkuss/iproto_server/internal/usecase"
	"github.com/kuzkuss/iproto_server/models"
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

	var res interface{}
	req, err := del.GetRequest(conn)
	if err == nil {
		res, err = del.HandleRequest(req)
		if err == nil {
			if err := del.SendResponse(conn, req.Header, res, models.OK); err != nil {
				log.Println(err)
				return
			}
			return
		}
	}

	if req != nil {
		if err := del.SendResponse(conn, req.Header, err.Error(), models.ErrorCodes[err]); err != nil {
			log.Println(err)
			return
		}
	}
}

