package models

// User Userモデル
type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Token string `db:"token" json:"token"`
	UID   string `db:"uid" json:"uid"`
}
