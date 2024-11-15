package dto

type SubmitAnswersRequest struct {
	Answers map[string]string `json:"answers"` // Key: Question ID, Value: User's Answer
}
