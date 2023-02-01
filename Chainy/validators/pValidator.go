package validators

import (
	"1000chain/wallet"
	"time"
)

type Validator struct {
	Wallet     *wallet.Wallet
	Since time.Time
}

func NewValidator(owned float64) (*Validator, error) {
	w := wallet.NewWallet()
			
	V := &Validator{
		Wallet: w,
		Since: time.Now(),	
	}
	w.AddStake(owned)

	return V, nil
}