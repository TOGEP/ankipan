package models

type Card struct {
  Id int `db:"id" json:"id"`
  Problem string `db:"problex_statement" json:"problem_statement"`
  Anser string `db:"answer_text" json:"answer_text"`
  Memo string `db:"memo" json:"memo"`
  SolvedCount int `db:"solved_count" json:"solved_count"`
}
