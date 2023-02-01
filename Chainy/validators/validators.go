package validators

import (
	"errors"
	"fmt"
	"sort"
	
)

type Validators struct {
	All         []*Validator
	BestTen     []*Validator
	BestHundred []*Validator
}

func CreateRandomValidators(n int) []*Validator {
	v := make([]*Validator, 0)
	for i := 0; i < n; i++ {
		val, _ := NewValidator(float64(i * i))
		if val.PubKey != "" {
			v = append(v, val)
		}
	}

	fmt.Println(v)
	return v
}

func (v *Validators) BiggestOwner() *Validator {
	var maxOwner *Validator
	for _, val := range v.All {
		if val.Stake > maxOwner.Stake {
			maxOwner = val
		}
	}
	return maxOwner
}

func (v *Validators) sort() {
	if !sort.SliceIsSorted(v.All, func(i, j int) bool { return v.All[i].Stake > v.All[j].Stake }) {
		sort.Slice(v.All, func(i, j int) bool { return v.All[i].Stake > v.All[j].Stake })

	}
}

func (v *Validators) AddValidator(vld *Validator) error {
	v.sort()
	for _, val := range v.All {
		if val.PubKey == vld.PubKey {
			return errors.New("Validator is already exist")
		}
	}
	v.All = append(v.All, vld)
	return nil
}
