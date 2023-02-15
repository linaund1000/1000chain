package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
)

func GeneratePUBKEY() string {
	//is it good approach to bind to time	
	
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)	
	time.Sleep( time.Duration(r) * time.Millisecond)	

	var (
		s = sha256.New()
		now    = strconv.Itoa(int(time.Now().UnixMilli()))
		_, err = s.Write([]byte(now))
	)
	if err != nil {
		return ""
	}	
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}


