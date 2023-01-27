package Block

import (
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	t time.Time

	PHash []byte
	Data  int
	Hash  []byte
}

func FirstBlock() *Block {

	data := "dlkfjas;"

	h := sha256.New()
	h.Write([]byte(data))

	hs := h.Sum(nil)
	b := &Block{
		PHash: hs,
		Data:  0,
		Hash:  []byte{},
	}

	b.t = time.Now()
	b.Hash = b.generateHash()
	return b
}

func (B *Block) generateHash() []byte {
	t := int(time.Now().UnixNano())

	record := strconv.Itoa(B.Data) + string(B.PHash) + strconv.Itoa(t)
	h := sha256.New()
	h.Write([]byte(record))

	return h.Sum(nil)
}

func NewBlock(prvBlock *Block, data int) *Block {
	b := &Block{
		PHash: prvBlock.Hash,
		Data:  data,
	}
	b.Hash = b.generateHash()
	return b
}
