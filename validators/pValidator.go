package validators

import (
	"encoding/base64"	
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
	"errors"
)

type Validator struct {
	PubKey string
	stake  float64
}

func NewValidator(owned float64) (*Validator ,error) {
	pubk := generatePUBKEY()
	fmt.Println("pubkey generated")
	owned = owned *2
	
	if pubk ==""{
		return &Validator{
			PubKey: "",
			stake: 0,
		} , errors.New("generate pubkey error ") 
	}
	V := &Validator{
		PubKey: pubk,
		stake:  owned,
	}
	return V , nil
}
func generatePUBKEY()string{

	var (
		s   = sha256.New()
		
		now = strconv.Itoa(int(time.Now().UnixMilli()))
		_, err = s.Write([]byte(now))
	)
	if err != nil{
		return ""
		
	}	
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}