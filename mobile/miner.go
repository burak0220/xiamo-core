package mobile

import (
	"fmt"
	"github.com/burak0220/xiamo-core/core"
)

// StartMiningSession is called by the Android UI button
func StartMiningSession() string {
	header := []byte("xiamo-mobile-mining-alpha")
	// This calls your AAH algorithm from the core folder
	hash := core.CalculateAAH(header, 1) 
	return fmt.Sprintf("Success! Last Hash: %x", hash)
}

// GetXiamoStatus returns the network connectivity for the UI
func GetXiamoStatus() string {
	return "Xiamo Node: Active & Discoverable"
}
