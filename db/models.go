package db

import "time"

type Todo struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	Name string `json:"name,omitempty"`
}
