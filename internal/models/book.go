package models

type Book struct {
	ID     int    `db:"id"    json:"id"`
	Name   string `db:"name"  json:"name"`
	Title  string `db:"title" json:"title"`
	Author string `db:"author" json:"author"`
}
