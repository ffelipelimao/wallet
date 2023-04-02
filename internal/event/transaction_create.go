package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload any
}

func NewTransactionCreated() *TransactionCreated {
	return &TransactionCreated{
		Name: "TransactionCreated",
	}
}

func (e *TransactionCreated) GetName() string {
	return e.Name
}

func (e *TransactionCreated) GetPayload() any {
	return e.Payload
}

func (e *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}

func (e *TransactionCreated) SetPayload(payload any) {
	e.Payload = payload
}
