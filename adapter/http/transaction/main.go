package transaction

import (
	"encoding/json"
	"finance-planner/model/transaction"
	"finance-planner/util"
	"fmt"
	"net/http"
)

// GetTransactions returns all the transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Transaction 1",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.DateTime(),
		},
	}

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

// CreateTransactions creates a new transaction
func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response transaction.Transactions
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(response)
}
