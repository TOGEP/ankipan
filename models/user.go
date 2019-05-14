package models

type User struct {
  Id int `db:"id" json:"id"`
  Name string `db:"name" json:"name"`
  Email string `db:"email" json:"email"`
  Token string `db:"token" json:"token"`
  Uid string `db:"uid" json:"uid"`
}
