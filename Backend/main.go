package main

import (
	"1000chain/core"
	b "1000chain/core/block"
	v "1000chain/core/validators"

	"github.com/labstack/echo/v4"
)






func main() {
	

	
	e := echo.New()	

	cor := core.NewCore()
 	
	
	
	// chain = append(chain, fBlock)	
	// for i , validator := range val{
	// 	fmt.Println("CRINGE ", validator)
	// 	fmt.Println("stake ::",  validator.Wallet.Stake)
	// 	new := b.NewBlock(fBlock,i)
	// 	fBlock = new
	// 	chain = append(chain, fBlock)
	// }



	
	e.Start("4000")

}
