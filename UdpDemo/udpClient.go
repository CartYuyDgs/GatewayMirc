package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127,0,0,1),
		Port: 9000,
	})

	if err != nil {
		fmt.Printf("conn failed, err :%v\n", err)
		return
	}

	for i := 0; i<100;i++ {
		_, err := conn.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("write failed, err :%v\n", err)
			return
		}
	}
}