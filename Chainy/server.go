package main
import(
	b	"1000chain/Chainy/Block"
	v  	"1000chain/Chainy/validators"
)
	

type Server struct{
	Chain []*b.Block
	Validators []*v.Validator

}


func NewServer()*Server{
	s := &Server{
		Chain: make([]*b.Block, 0),		
		Validators: make([]*v.Validator,0),	
		
		// redis cache 
		// 
	}
	return s
}
