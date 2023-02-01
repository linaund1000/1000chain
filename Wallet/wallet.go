package wallet

import (
	c "1000chain/crypto"
	"time"
)
type Record struct {
	IntoWallet map[time.Time]float64

}



type Wallet struct {	
	PubKey string
	Stake  float64
}


func NewWallet()*Wallet{
	w := &Wallet{
		PubKey: "",
		Stake : 0 ,
	}
	w.PubKey = c.GeneratePUBKEY()

	return	w
}

func (w *Wallet)AddStake(amount float64){

	w.Stake = amount + w.Stake
}