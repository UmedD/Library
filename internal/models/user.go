package models

type User struct {
	ID       int    `db:"id" json:"id,omitempty"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"-"`
}
