package repository

import "quiz-app/server/model"

type LeaderboardRepository interface {
	SaveScore(username string, score int)
	GetLeaderboard() []model.LeaderboardEntry
	GetAllScores() []int
}
