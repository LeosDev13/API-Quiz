package dto

type LeaderboardEntry struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
	Rank     int    `json:"rank"`
}

type LeaderboardResponse struct {
	Entries []LeaderboardEntry `json:"entries"`
}
