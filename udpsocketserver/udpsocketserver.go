package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	resolver := "udp"
	serverAddr := "127.0.0.1:1234"
	udpAddr, err := net.ResolveUDPAddr(resolver, serverAddr)
	if err != nil {
		log.Fatalf("ResolveUDPAddr failed: %s", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("estanblish listerner failed: %s\n", err)
	}

	for {
		fmt.Println("receive msg")
		buf := make([]byte, 512)
		_, remoteaddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalf("receive msg failed: %s\n", err)
		}
		fmt.Printf("receive from %s, %s", remoteaddr, buf)

		buf = []byte(remoteaddr.String())
		_, err = conn.WriteToUDP(buf[:], remoteaddr)
		if err != nil {
			log.Fatalf("send msg failed: %s", err)
		}
	}

}
