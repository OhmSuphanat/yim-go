package entity

import "time"

type Book struct {
	Id	   int64  `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Author   string `db:"author" json:"author"`
	Description string `db:"description" json:"description"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}