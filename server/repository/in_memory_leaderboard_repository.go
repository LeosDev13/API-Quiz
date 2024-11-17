package repository

import (
	"quiz-app/shared/dto"
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

func (repo *InMemoryLeaderboardRepository) GetLeaderboard() []dto.LeaderboardEntry {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	var leaderboard []dto.LeaderboardEntry
	for username, score := range repo.scores {
		leaderboard = append(leaderboard, dto.LeaderboardEntry{
			Username: username,
			Score:    score,
		})
	}

	return leaderboard
}
