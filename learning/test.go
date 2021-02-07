package main

import (
	"fmt"
	"net"
)

func main() {

	addrs, _ := net.InterfaceAddrs()
	fmt.Println(addrs)
}
