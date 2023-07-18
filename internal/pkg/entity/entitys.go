package entity

import "time"

type (
	User struct {
		ID        uint64    `json:"id"`
		Balance   float64   `json:"balance"`
		CreatedDT time.Time `json:"created_dt"`
	}

	Operation struct {
		ID            uint64  `json:"id"`
		OperationType string  `json:"operation_type"`
		Amount        float64 `amount:"amount"`
		Currency      string  `json:"currency"`
		CreatedDT     int64   `json:"created_dt"`
	}
)
