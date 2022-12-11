package dto

type NewAccountResponse struct {
	AccountId string `json:"id"`
}

type AccountResponse struct {
	AccountId   string `json:"id"`
	CustomerId  string `json:"customer_id"`
	Amount      string `json:"ammount"`
	OpeningDate string `json:"opening_date"`
	AccountType string `json:"account_type"`
	Status      string `json:"status"`
}
