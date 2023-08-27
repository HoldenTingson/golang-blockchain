package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"goblockchain/utils"
)

type Transaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	senderAddress    string
	recipientAddress string
	value            float32
}

type TransactionRequest struct {
	SenderPrivateKey           *string `json:"sender_private_key"`
	SenderBlockchainAddress    *string `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string `json:"recipient_blockchain_address"`
	SenderPublicKey            *string `json:"sender_public_key"`
	Value                      *string `json:"value"`
}

func (tr *TransactionRequest) Validate() bool {
	if *tr.SenderPrivateKey == "" ||
		*tr.SenderBlockchainAddress == "" ||
		*tr.RecipientBlockchainAddress == "" ||
		*tr.SenderPublicKey == "" ||
		*tr.Value == "" {

		return false
	}
	fmt.Println(*tr.RecipientBlockchainAddress)
	return true
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *Transaction {
	return &Transaction{
		senderPrivateKey: privateKey,
		senderPublicKey:  publicKey,
		senderAddress:    sender,
		recipientAddress: recipient,
		value:            value,
	}
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &utils.Signature{
		R: r,
		S: s,
	}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderAddress,
		Recipient: t.recipientAddress,
		Value:     t.value,
	})
}
