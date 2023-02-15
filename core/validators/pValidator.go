package validators

import (
	"1000chain/wallet"
	"1000chain/crypto"
	"time"
)

type Validator struct {
	Wallet     *wallet.Wallet
	Since time.Time
	Id string 
}

func NewValidator(owned float64) (*Validator, error) {
	w := wallet.NewWallet()
			
	V := &Validator{
		Wallet: w,
		Since: time.Now(),		
	}
	w.AddStake(owned)
	V.Id = crypto.GeneratePUBKEY()
	return V, nil
}

