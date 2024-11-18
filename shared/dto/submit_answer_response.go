package dto

type SubmitAnswersResponse struct {
	CorrectAnswers int     `json:"correct_answers"`
	Percentile     float64 `json:"percentile"`
}
