package handler

import (
	"encoding/json"
	"net/http"
	"quiz-app/dto"
	"quiz-app/repository"
)

type AnswerHandler struct {
	questionRepo repository.QuestionRepository
}

func NewSubmitAnswersHandler(repo repository.QuestionRepository) *AnswerHandler {
	return &AnswerHandler{
		questionRepo: repo,
	}
}

func (h *AnswerHandler) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var req dto.SubmitAnswersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	correctAnswers := 0
	questions := h.questionRepo.GetQuestions()

	for _, question := range questions {
		if userAnswer, ok := req.Answers[question.ID]; ok && userAnswer == question.Answer {
			correctAnswers++
		}
	}

	response := dto.SubmitAnswersResponse{
		CorrectAnswers: correctAnswers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
