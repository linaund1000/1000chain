package main

import (
	b "1000chain/Block"
	v "1000chain/validators"
	"fmt"
	"time"
)

func main() {
	chain := make([]*b.Block, 0)
	start := time.Now()
 	val := v.CreateRandomValidators(100)
	fBlock := b.FirstBlock()
	chain = append(chain, fBlock)	
	for i , validator := range val{
		fmt.Println("CRINGE ", validator.PubKey)
		fmt.Println("                        ")
		new := b.NewBlock(fBlock,i)
		fBlock = new
		chain = append(chain, fBlock)
	}

	fmt.Println(val)


	
	
	
	
	// for i := 0; i < 1000000; i++ {
	// 		blckx := b.NewBlock(prvBlck, i)
	// 		chain = append(chain, blckx)
	// 		prvBlck = blckx
	// }
	diff := time.Since(start)
	fmt.Println(diff)
}
