package core

import "github.com/burak0220/xiamo-core/network"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// GenerateMobileWallet creates a unique address for the user
func GenerateMobileWallet() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	fmt.Printf("[WALLET] New Address Generated: %x\n", privateKey.PublicKey)
}
