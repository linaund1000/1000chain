package wallet

import (
	c "1000chain/crypto"
	"fmt"
	"time"
)

type Record struct {
	IntoWallet map[int64]float64
	FromWallet map[int64]float64
}

func (r *Record) Into(amound float64) {
	time := time.Now().UnixNano()
	r.IntoWallet[time] = amound
}
func (r *Record) From(amound float64) {
	time := time.Now().UnixNano()
	r.FromWallet[time] = amound
}

type Wallet struct {
	*Record
	PubKey string
	Stake  float64
}

func NewWallet() *Wallet {
	w := &Wallet{
		Record: &Record{},
		PubKey: "",
		Stake:  0,
	}
	w.PubKey = c.GeneratePUBKEY()
	return w
}

func (w *Wallet) AddStake(amount float64) error {
	w.Stake = amount + w.Stake
	w.Into(amount)
	return nil
}
func (w *Wallet) SubtractStake(amount float64) error {
	if w.Stake < amount {
		return	fmt.Errorf("not enough money in stake" ) 
	}

	w.Stake -= amount
	w.From(amount)
	return nil
}
