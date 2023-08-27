package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	timestamp   int64
	nonce       int
	prevHash    [32]byte
	transaction []*Transactions
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transactions) *Block {
	return &Block{
		timestamp:   time.Now().UnixNano(),
		nonce:       nonce,
		prevHash:    prevHash,
		transaction: transactions,
	}
}

func (b *Block) Hash() [32]byte {
	a, _ := json.Marshal(b)
	return sha256.Sum256([]byte(a))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64           `json:"timestamp"`
		Nonce        int             `json:"nonce"`
		PrevHash     string          `json:"previous_hash"`
		Transactions []*Transactions `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PrevHash:     fmt.Sprintf("%x", b.prevHash),
		Transactions: b.transaction,
	})

}

func (b *Block) Print() {
	fmt.Printf("timestamp	%d\n", b.timestamp)
	fmt.Printf("nonce		%d\n", b.nonce)
	fmt.Printf("prevHash	%x\n", b.prevHash)
	for _, t := range b.transaction {
		t.PrintTransaction()
	}
}
