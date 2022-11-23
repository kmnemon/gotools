package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	resolver := "tcp"
	serverAddr := "127.0.0.1:1234"
	tcpAddr, err := net.ResolveTCPAddr(resolver, serverAddr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed: %s", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("estanblish listerner failed: %s\n", err)
	}

	fmt.Println("waiting connection")
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("listerner accept failed: %s\n", err)
	}
	fmt.Println("connected")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		fmt.Println("send msg")
		if _, err := conn.Write([]byte(text)); err != nil {
			log.Fatalf("send msg failed: %s\n", err)
		}

		fmt.Println("receive msg")
		buf := make([]byte, 512)
		_, err := conn.Read(buf[0:])
		if err != nil {
			log.Fatalf("receive msg failed: %s\n", err)
		}

		fmt.Printf("print msg: %s\n", buf)
	}

}
