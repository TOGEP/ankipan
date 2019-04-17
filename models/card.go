package models

type Card struct {
  Id int `db:"id" json:"id"`
  Problem string `db:"id" json:"problem"`
  Anser string `db:"anser" json:"anser"`
  Note string `db:"note" json:"note"`
}

/*
type InputObjectConfig struct {
    Name        string      `json:"name"`
    Fields      interface{} `json:"fields"`
    Description string      `json:"description"`
}*/

