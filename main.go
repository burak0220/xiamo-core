package main

import (
	"fmt"
	"github.com/burak0220/xiamo-core/core"
	"time"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("   XIAMO NETWORK (XMO) - MOBILE L1      ")
	fmt.Println("   Version: v0.1.0-Alpha                ")
	fmt.Println("========================================")

	// 1. Initialize P2P Networking
	core.StartP2P("1611")

	// 2. Mining Parameters
	header := []byte("xiamo-genesis-block-2026")
	var nonce uint64 = 0

	fmt.Println("\n[*] Node Status: Online")
	fmt.Println("[*] Mining Policy: 1 Block-per-Second (1bps)")
	fmt.Println("[*] Algorithm: ARM-Adaptive HeavyHash (AAH)")
	fmt.Println("----------------------------------------")

	// 3. Main Consensus Loop
	for {
		// Calling your custom algorithm from core/pow_engine.go
		hash := core.CalculateAAH(header, nonce)
		
		fmt.Printf("[BLOCK MINED] Nonce: %d | Hash: %x\n", nonce, hash)
		
		nonce++
		
		// Target 1bps synchronization
		time.Sleep(time.Second)
	}
}
