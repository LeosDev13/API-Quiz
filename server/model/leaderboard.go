package model

type LeaderboardEntry struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
	Rank     int    `json:"rank"`
}
