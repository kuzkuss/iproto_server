package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/kuzkuss/iproto_server/models"
	"github.com/vmihailenco/msgpack/v5"
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

	request := models.Request {
		Header: models.Header{
			FuncId: 0x00020001,
			RequestId: 1,
		},
	}

	b, _ := msgpack.Marshal(models.RequestSaveString {
		Idx: 1,
		Str: "hello",
	})

	request.Body = b
	request.Header.BodyLength = uint32(len(request.Body))

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, &request.Header); err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := conn.Write(buf.Bytes()); err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, err := conn.Write(request.Body); err != nil {
		fmt.Println(err.Error())
		return
	}

	headerBytes := make([]byte, models.SIZE_UINT * 3)
	_, err = conn.Read(headerBytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp := models.Response{}
	headerBuf := bytes.NewBuffer(headerBytes)

	if err := binary.Read(headerBuf, binary.LittleEndian, &resp.Header.FuncId); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := binary.Read(headerBuf, binary.LittleEndian, &resp.Header.BodyLength); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := binary.Read(headerBuf, binary.LittleEndian, &resp.Header.RequestId); err != nil {
		fmt.Println(err.Error())
		return
	}

	codeBytes := make([]byte, models.SIZE_UINT)
	_, err = conn.Read(codeBytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	codeBuf := bytes.NewBuffer(codeBytes)

	if err := binary.Read(codeBuf, binary.LittleEndian, &resp.ReturnCode); err != nil {
		fmt.Println(err.Error())
		return
	}

	resp.Body = make([]byte, resp.Header.BodyLength)

	_, err = conn.Read(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var str string

	if err := msgpack.Unmarshal(resp.Body, &str); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(str)
}

