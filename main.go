package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to connect with the port", err.Error())
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error Accepting connection", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	//incoming request started
	buf := make([]byte, 1024)

	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error while reading the message", err.Error())
		os.Exit(1)
	}
	fmt.Printf("received message %v (%d)", buf[8:12], int32(binary.BigEndian.Uint32(buf[8:12])))
	var version_error []byte
	ver := binary.BigEndian.Uint16(buf[4:6])
	// fmt.Printf("received message, %v, %v", ver, buf[4:8])
	switch ver {
	case 0, 1, 2, 3, 4:
		version_error = []byte{0, 0}
	default:
		version_error = []byte{0, 35}

	}

	// response
	resp := make([]byte, 8)
	copy(resp, []byte{0, 0, 0, 0})
	copy(resp[4:], buf[8:12])
	resp = append(resp, version_error...)
	conn.Write(resp)

}
