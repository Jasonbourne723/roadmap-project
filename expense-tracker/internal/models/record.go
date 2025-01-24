package models

import "time"

type Record struct {
	Id          int32     `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
}
