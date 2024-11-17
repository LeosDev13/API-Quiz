package dto

type SubmitAnswersRequest struct {
	Username string            `json:"username"`
	Answers  map[string]string `json:"answers"`
}
