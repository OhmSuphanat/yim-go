package domain

import "time"

type CreateBookRequest struct {
	Title    string
	Author   string
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}