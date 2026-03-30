package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Payment struct {
	ID            uint
	OrderID       string
	TransactionID string
	Amount        int64  // amount in cents
	Status        string // Authorized/Declined
}

func GenerateTransactionID(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hashBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashBytes)
}
