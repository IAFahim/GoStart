package main

import (
	"net"
)

//type P2PServer struct {
//	p2pPort  int
//	listener *net.UDPConn
//}

func main() {
	pc, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 9806})
	go HandlePacket(pc)
	select {}
}

func HandlePacket(pc *net.UDPConn) {
	buf := make([]byte, 528)
	for {
		size, addr, err := pc.ReadFromUDP(buf)
		if err != nil {
			break
		}
		data := buf[:size]
		println(string(data))
		_, _ = pc.WriteTo(append([]byte("hello me"), data...), addr)
	}
	_ = pc.Close()
}
