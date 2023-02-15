package validators

import (
	"errors"
	"sort"
	// "fmt"
)

type Validators struct {
	All    []*Validator
	AllIds []string
	// BestTen     []*Validator / []string 
	// BestHundred []*Validator
}

func NewValidatorsSlice() *Validators {
	v := &Validators{
		All:    make([]*Validator, 0),
		AllIds: make([]string, 0),
	}
	return v
}


// ... i dont think we can add multiple validator at the same time 
// func CreateRandomValidators(n int) []*Validator {
// 	v := make([]*Validator, 0)
// 	for i := 0; i < n; i++ {
// 		val, _ := NewValidator(float64(i * i))
// 		if val.Id != "" {
// 			v = append(v, val)
// 		}
// 	}

// 	// fmt.Println(v)
// 	return v
// }

func (v *Validators) BiggestOwner() *Validator {
	var maxOwner *Validator
	for _, val := range v.All {
		if val.Wallet.Stake > maxOwner.Wallet.Stake {
			maxOwner = val
		}
	}
	return maxOwner
}

func (v *Validators) sort() {
	if !sort.SliceIsSorted(v.All, func(i, j int) bool { return v.All[i].Wallet.Stake > v.All[j].Wallet.Stake }) {
		sort.Slice(v.All, func(i, j int) bool { return v.All[i].Wallet.Stake > v.All[j].Wallet.Stake })

	}
}

func (v *Validators) AddValidator(vld *Validator) error {
	// v.sort()
	for _, val := range v.All {
		if val.Id == vld.Id {
			return errors.New("Validator is already exist")
		}
	}

	v.All = append(v.All, vld)
	v.AllIds = append(v.AllIds, vld.Id)
	return nil
}
