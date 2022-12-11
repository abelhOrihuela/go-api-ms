package dto

import (
	"strings"

	"banking.com/abelh/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	Amount      float64 `json:"amount"`
	AccountType string  `json:"account_type"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
