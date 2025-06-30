package models

type Book struct {
	ID         int    `db:"id"       json:"id"`
	Name       string `db:"name"     json:"name"`
	Title      string `db:"title"    json:"title"`
	AuthorID   int    `db:"author_id" json:"author_id"`
	AuthorName string `db:"author_name" json:"author"`
}
