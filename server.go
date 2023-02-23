package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"net"
)

const (
	ADM_STORAGE_SWITCH_READONLY = 0x00010001
	ADM_STORAGE_SWITCH_READWRITE = 0x00010002
	ADM_STORAGE_SWITCH_MAINTENANCE = 0x00010003
	STORAGE_REPLACE = 0x00020001
	STORAGE_READ = 0x00020002
)

type Header struct {
	funcId uint32
	bodyLength uint32
	requestId uint32
}

type Request struct {
	header Header
	body []byte
}

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

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn) {
	defer conn.Close()
	req, err := getRequest(conn)
	if err != nil {
		log.Println(err)
		return
	}

	err = handleRequest(req)
	if err != nil {
		log.Println(err)
		return
	}

	if err := sendResponse(conn); err != nil {
		log.Println(err)
		return
	}
}

func getRequest(conn *net.TCPConn) (*Request, error) {
	headerBytes := make([]byte, 12)
	_, err := conn.Read(headerBytes)
	if err != nil {
		return nil, err
	}

	req := Request{}
	headerBuf := bytes.NewBuffer(headerBytes)

	if err := binary.Read(headerBuf, binary.LittleEndian, &req.header.funcId); err != nil {
		return nil, err
	}

	if err := binary.Read(headerBuf, binary.LittleEndian, &req.header.bodyLength); err != nil {
		return nil, err
	}

	if err := binary.Read(headerBuf, binary.LittleEndian, &req.header.requestId); err != nil {
		return nil, err
	}

	req.body = make([]byte, req.header.bodyLength)

	_, err = conn.Read(req.body)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func handleRequest(request *Request) error {
	switch request.header.funcId {
	case ADM_STORAGE_SWITCH_READONLY:
		switchReadOnly(request)

	case ADM_STORAGE_SWITCH_READWRITE:
	
	case ADM_STORAGE_SWITCH_MAINTENANCE:
		
	case STORAGE_REPLACE:
		
	case STORAGE_READ:
	
	default:
		return errors.New("invalid request func")
	}

	return nil
}

func switchReadOnly(request *Request) {

}

