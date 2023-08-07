package entity

import "time"

type (
	User struct {
		ID        uint64    `json:"id,omitempty"`
		Balance   float64   `json:"balance"`
		Currency  string    `json:"currency,omitempty"`
		CreatedDT time.Time `json:"created_dt,omitempty"`
	}

	Operation struct {
		ID            uint64    `json:"id,omitempty"`
		OperationType string    `json:"operation_type,omitempty"`
		Amount        float64   `json:"amount,omitempty"`
		Currency      string    `json:"currency,omitempty"`
		CreatedDT     time.Time `json:"created_dt,omitempty"`
	}
)
