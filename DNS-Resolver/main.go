package main

import (
	"dnsresolver/internal/model"
	"fmt"
	"net"
)

func main() {
	msg := model.NewQuestion(22, "dns.google.com")
	fmt.Printf("%x\n", msg.Encode())
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	wrote, _ := conn.Write(msg.Encode())
	fmt.Printf("Wrote %d\n", wrote)

	response := make([]byte, 65)
	read, err := conn.Read(response)
	fmt.Printf("Read %d\n", read)
	fmt.Printf("%x\n", response)
	conn.Close()
	responseMsg := model.ParseResponse(response)
	fmt.Println(responseMsg.Print())
}
