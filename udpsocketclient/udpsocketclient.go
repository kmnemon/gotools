package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	resolver := "udp"
	serverAddr := "127.0.0.1:1234"
	udpAddr, err := net.ResolveUDPAddr(resolver, serverAddr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed: %s", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Dial udp failed: %s", err)
	}

	for {
		//send msg
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text)

		//receive msg
		buf := make([]byte, 512)
		_, _, err = conn.ReadFrom(buf)
		if err != nil {
			log.Fatalf("server receive msg failed: %s", err)
		}
		fmt.Printf("client receive msg: %s\n", buf)
	}
}
