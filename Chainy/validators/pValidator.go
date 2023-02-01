package validators

import (
	c "1000chain/crypto"
	"1000chain/wallet"
	"errors"
	"fmt"
	"time"
)

type Validator struct {
	Wallet     wallet.Wallet
	Since time.Time
}

func NewValidator(owned float64) (*Validator, error) {
	pubk := c.GeneratePUBKEY()
	fmt.Println("pubkey generated")
	owned = owned * 2

	if pubk == "" {
		return &Validator{
				Wallet: wallet.Wallet{
					PubKey: "",
					Stake:  0,
				},
				Since: time.Now(),
			},
			errors.New("generate pubkey error ")
	}
	V := &Validator{
		Wallet: wallet.Wallet{
			PubKey: pubk,
			Stake:  owned,
		},
		Since: time.Now()}
	return V, nil
}