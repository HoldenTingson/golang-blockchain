package block

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transactions struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

type TransactionRequest struct {
	SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
	SenderPublicKey            *string  `json:"sender_public_key"`
	Value                      *float32 `json:"value"`
	Signature                  *string  `json:"signature"`
}

type AmountResponse struct {
	Amount float32 `json:"amount"`
}

func (ar *AmountResponse) MarshallJson() ([]byte, error) {
	return json.Marshal(struct {
		Amount float32 `json:"amount"`
	}{
		Amount: ar.Amount,
	})
}

func (tr *TransactionRequest) Validate() bool {
	if tr.Value == nil ||
		tr.SenderBlockchainAddress == nil ||
		tr.RecipientBlockchainAddress == nil ||
		tr.SenderPublicKey == nil ||
		tr.Signature == nil {

		return false
	}

	return true
}

func NewTransaction(sender string, recipient string, value float32) *Transactions {
	return &Transactions{
		senderAddress:    sender,
		recipientAddress: recipient,
		value:            value}
}

func (t *Transactions) PrintTransaction() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender address		%s\n", t.senderAddress)
	fmt.Printf("recipient address	%s\n", t.recipientAddress)
	fmt.Printf("sender address		%.1f\n", t.value)
}

func (t *Transactions) MarshalJSON() ([]byte, error) {
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
