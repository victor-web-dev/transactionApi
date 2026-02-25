package transaction

import "time"

// Transaction represents the data received or sent from a transaction
//
// Title purpose of the transaction
// Amount value related to the purpose
// Type describes the types of transaction,
// according to the numbers starting from 0
// CreatedAt exact time of the transaction
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions slice of Transaction for multiple transactions
type Transactions []Transaction
