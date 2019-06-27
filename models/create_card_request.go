package models

// CreateCardRequest request jsonをマッピングするためのモデル
type CreateCardRequest struct {
	ID          int    `json:"id"`
	Problem     string `json:"problem_statement"`
	Anser       string `json:"answer_text"`
	Memo        string `json:"memo"`
	SolvedCount int    `json:"solved_count"`
	Token       string `json:"token"`
}
