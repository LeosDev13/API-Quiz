package handler

import (
	"encoding/json"
	"net/http"
	"quiz-app/server/logger"
	"quiz-app/server/repository"
	"quiz-app/shared/dto"
	"sort"
)

type LeaderboardHandler struct {
	leaderboardRepo repository.LeaderboardRepository
	logger          logger.Logger
}

func NewLeaderboardHandler(repo repository.LeaderboardRepository, log logger.Logger) *LeaderboardHandler {
	return &LeaderboardHandler{
		leaderboardRepo: repo,
		logger:          log,
	}
}

func (h *LeaderboardHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	leaderboard := h.leaderboardRepo.GetLeaderboard()

	sort.SliceStable(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	for i := range leaderboard {
		leaderboard[i].Rank = i + 1
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.LeaderboardResponse{Entries: leaderboard}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		h.logger.Error(err, "Failed to encode leaderboard response")
		return
	}
}
