package validators_test

import (
	validators "1000chain/alidators"
	"fmt"
	"testing"
	//"github.com/stretchr/testify/assert"
)
func TestNewValidators(t *testing.T) {
	vals := validators.CreateRandomValidators(100)	
	if len(vals)>100{
		t.Errorf("error")
	}
	
	fmt.Println(vals)
}