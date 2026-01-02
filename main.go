package main

import (
    "fmt"
    "github.com/burak0220/xiamo-core/core"
    "time"
)

func main() {
    fmt.Println("Xiamo Network L1 - Mobile BlockDAG Starting...")
    fmt.Println("Mining started at 1 block-per-second...")

    for i := 0; ; i++ {
        fmt.Printf("Block #%d Mined!\n", i)
        time.Sleep(time.Second)
    }
}
