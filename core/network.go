package core

import (
	"fmt"
	"net"
)

// XiamoNode represents a mobile peer in the network
type XiamoNode struct {
	ID      string
	Address string
	Peers   []string
}

// StartP2P initializes the mobile mesh network
func StartP2P(port string) {
	fmt.Printf("Xiamo P2P Layer starting on port %s...\n", port)
	fmt.Println("Searching for other mobile nodes via Discovery...")
	
	// Bu kısım diğer telefonların IP adreslerini bulacak
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting network:", err)
		return
	}
	fmt.Println("Node is now ONLINE and discoverable.")
	_ = ln
}
