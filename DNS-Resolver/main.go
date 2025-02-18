package main

import (
	"dnsresolver/internal/model"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No argumensts")
		return
	}
	domainName := os.Args[1]

	msg := model.NewQuestion(uint16(rand.Int64N(1000)), domainName)
	ips, err := msg.ResolveName()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	if ips == nil {
		fmt.Println("No ips found...")
		return
	}
	fmt.Printf("Ips of the: %s\n", domainName)
	for i, ip := range ips {
		fmt.Printf("[%d]: %s\n", i, ip)
	}
}
