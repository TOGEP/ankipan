package models

type Card struct {
  Id int `db:"id" json:"id"`
  Problem string `db:"problex_statement" json:"problem_statement"`
  Anser string `db:"anser_text" json:"anser_text"`
  Note string `db:"memo" json:"memo"`
}
