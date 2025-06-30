package models

type Author struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Books []Book `json:"books,omitempty"`
}
