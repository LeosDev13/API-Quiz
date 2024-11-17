package handler

import (
	"encoding/json"
	"net/http"
	"quiz-app/server/model"
	"quiz-app/server/repository"
	"quiz-app/shared/dto"
)

func mapToDTO(q model.Question) dto.Question {
	return dto.Question{
		ID:       q.ID,
		Question: q.Question,
		Options:  q.Options,
	}
}

type QuestionHandler struct {
	questionRepo repository.QuestionRepository
}

func NewQuestionHandler(repo repository.QuestionRepository) *QuestionHandler {
	return &QuestionHandler{
		questionRepo: repo,
	}
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions := h.questionRepo.GetQuestions()

	var questionDTOs []dto.Question
	for _, q := range questions {
		questionDTOs = append(questionDTOs, mapToDTO(q))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(questionDTOs); err != nil {
		http.Error(w, "Failed to encode questions", http.StatusInternalServerError)
		return
	}
}
