package main

import (
	b "1000chain/block"
	"fmt"
	"time"
)

func main() {
	chain := make([]*b.Block, 0)

	bck := b.FirstBlock()
	start := time.Now()
	prvBlck := bck
	for i := 0; i < 1000000; i++ {
			blckx := b.NewBlock(prvBlck, i)
			chain = append(chain, blckx)
			prvBlck = blckx
	}
	diff := time.Since(start)
	fmt.Println(diff)
}
