package core

import "github.com/burak0220/xiamo-core/network"
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
	fmt.Printf("[P2P] Starting Xiamo Network Layer on port %s...\n", port)
	
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("[ERROR] Network initialization failed: %v\n", err)
		return
	}
	fmt.Println("[P2P] Node is now listening for inbound connections.")
	_ = ln
}
