package repository

import (
	"quiz-app/shared/dto"
)

type LeaderboardRepository interface {
	SaveScore(username string, score int)
	GetLeaderboard() []dto.LeaderboardEntry
}
