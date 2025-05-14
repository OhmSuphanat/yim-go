package domain

import "time"

type Book struct {
	Id	   int64  
	Title    string
	Author   string
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}