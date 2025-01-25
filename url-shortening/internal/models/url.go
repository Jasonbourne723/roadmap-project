package models

import "time"

type Url struct {
	Id        int
	Short     string
	Long      string
	CreatedAt time.Time
}
