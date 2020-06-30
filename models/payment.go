package models

// GetTransactionMutationRequest func
type GetTransactionMutationRequest struct {
	ID            string  `json:"id"`
	BankID        string  `json:"bank_id"`
	AccountNumber string  `json:"account_number"`
	BankType      string  `json:"bank_type"`
	Date          string  `json:"date"`
	Amount        int     `json:"amount"`
	Description   string  `json:"description"`
	Type          string  `json:"type"`
	Balance       float64 `json:"balance"`
}

type DepositParameter struct {
	AccountName   string  `json:"accountName"`
	AccountNumber int     `json:"accountNumber"`
	Amount        float64 `json:"amount"`
}

type ResponseDeposit struct {
	Amount          int    `json:"amount"`
	TransferAccount string `json:"transferAccount"`
}
