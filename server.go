package main
import(
	"net"
	b	"1000chain/Block"
	v  	"1000chain/validators"
)
	

type Server struct{
	chain []*b.Block
	validators []*v.Validator

}


func NewServer()*Server{
	s := &Server{
		chain: make([]*b.Block, 0),		
		validators: []*v.Validator,	
		
		// redis cache 
		// 
	}
	return s
}
