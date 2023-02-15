package core

import (
	b "1000chain/core/block"
	v "1000chain/core/validators"
)

type Core struct {
	Chain      []b.Block
	Validators v.Validators
	// Validators []v.Validator.Id
	Fblock b.Block
}

func NewCore() *Core {
	c := &Core{
		Chain:      make([]b.Block, 0),
		Fblock:     *b.FirstBlock(),
	}
	c.Validators = v.CreateRandomValidators(10)
	return c
}

func (c *Core) AddValidator(v v.Validator) error {
	//we need to validate validators with a database query but for now we just accept them
	c.Validators = append(c.Validators, v.Id)
}
