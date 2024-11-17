package handler

import (
	"encoding/json"
	"net/http"
	"quiz-app/server/logger"
	"quiz-app/server/repository"
	"quiz-app/shared/dto"
)

type AnswerHandler struct {
	questionRepo    repository.QuestionRepository
	leaderboardRepo repository.LeaderboardRepository
	logger          logger.Logger
}

func NewSubmitAnswersHandler(
	questionRepo repository.QuestionRepository,
	leaderboardRepo repository.LeaderboardRepository,
	log logger.Logger,
) *AnswerHandler {
	return &AnswerHandler{
		questionRepo:    questionRepo,
		leaderboardRepo: leaderboardRepo,
		logger:          log,
	}
}

func (h *AnswerHandler) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var req dto.SubmitAnswersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		h.logger.Error(err, "Failed to decode request payload")
		return
	}

	correctAnswers := 0
	questions := h.questionRepo.GetQuestions()

	for _, question := range questions {
		if userAnswer, ok := req.Answers[question.ID]; ok && userAnswer == question.Answer {
			correctAnswers++
		}
	}

	h.leaderboardRepo.SaveScore(req.Username, correctAnswers)

	response := dto.SubmitAnswersResponse{
		CorrectAnswers: correctAnswers,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
