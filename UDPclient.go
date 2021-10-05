package main

import "net"

func main() {
	pc, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{103, 84, 253, 122}, Port: 9806})
	SendPacket(pc)
}

func SendPacket(pc *net.UDPConn) error {
	defer pc.Close()
	pc.Write([]byte("hello"))
	return nil
}
