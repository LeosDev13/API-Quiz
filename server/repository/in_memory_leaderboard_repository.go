package repository

import (
	"quiz-app/server/model"
	"sync"
)

type InMemoryLeaderboardRepository struct {
	scores map[string]int
	mu     sync.RWMutex
}

func NewInMemoryLeaderboardRepository() *InMemoryLeaderboardRepository {
	return &InMemoryLeaderboardRepository{
		scores: make(map[string]int),
	}
}

func (repo *InMemoryLeaderboardRepository) SaveScore(username string, score int) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.scores[username] = score
}

func (repo *InMemoryLeaderboardRepository) GetLeaderboard() []model.LeaderboardEntry {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var leaderboard []model.LeaderboardEntry
	for username, score := range repo.scores {
		leaderboard = append(leaderboard, model.LeaderboardEntry{
			Username: username,
			Score:    score,
		})
	}

	return leaderboard
}

func (repo *InMemoryLeaderboardRepository) GetAllScores() []int {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var scores []int
	for _, score := range repo.scores {
		scores = append(scores, score)
	}

	return scores
}
