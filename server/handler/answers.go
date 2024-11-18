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

	allScores := h.leaderboardRepo.GetAllScores()
	percentile := h.calculatePercentile(allScores, correctAnswers)

	response := dto.SubmitAnswersResponse{
		CorrectAnswers: correctAnswers,
		Percentile:     percentile,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *AnswerHandler) calculatePercentile(allScores []int, userScore int) float64 {
	countLowerScores := 0
	for _, score := range allScores {
		if score < userScore {
			countLowerScores++
		}
	}

	percentile := float64(countLowerScores) / float64(len(allScores)) * 100
	return percentile

}
