package main

import (
	"fmt"
	"log"
	"net"
)

func tcpclient() {
	resolver := "tcp"
	serverAddr := "127.0.0.1:1234"
	tcpAddr, err := net.ResolveTCPAddr(resolver, serverAddr)
	if err != nil {
		log.Fatalf("Client ResolveTCPAddr failed: %s", err)
	}

	conn, err := net.DialTCP(resolver, nil, tcpAddr)
	if err != nil {
		log.Fatalf("client DialTcp failed: %s", err)
	}

	for {
		var buf [512]byte
		_, err = conn.Read(buf[0:])
		if err != nil {
			log.Fatalf("client recevie msg failed: %s", err)
		}
		fmt.Printf("receive msg: %s\n", buf)

		sendMsg(conn)
	}

}

func sendMsg(conn *net.TCPConn) {

	message := []byte{'1', '2', '3'}
	_, err := conn.Write(message)
	if err != nil {
		log.Fatalf("client send msg failed: %s", err)
	}
}
