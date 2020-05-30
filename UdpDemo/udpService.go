package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp",&net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port: 9000,
	})

	if err != nil {
		fmt.Printf("listen failed, err :%v\n", err)
		return
	}

	for {
		var data [1024]byte

		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed, err :%v\n", err)
			return
		}

		fmt.Printf("n: %v, addr :%v , data: %s\n", n, addr, string(data[:]))
	}
}
