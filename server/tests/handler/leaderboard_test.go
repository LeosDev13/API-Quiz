package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quiz-app/server/handler"
	"quiz-app/server/logger"
	"quiz-app/server/repository"
	"quiz-app/shared/dto"
	"testing"
)

func TestGetLeaderboard(t *testing.T) {
	leaderboardRepo := repository.NewInMemoryLeaderboardRepository()

	leaderboardRepo.SaveScore("alice", 10)
	leaderboardRepo.SaveScore("bob", 20)
	leaderboardRepo.SaveScore("carol", 15)

	log := logger.New()
	handler := handler.NewLeaderboardHandler(leaderboardRepo, log)

	req := httptest.NewRequest("GET", "/leaderboard", nil)

	rr := httptest.NewRecorder()

	handler.GetLeaderboard(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %v", status)
	}

	var response dto.LeaderboardResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(response.Entries) != 3 {
		t.Fatalf("Expected 3 leaderboard entries, got %d", len(response.Entries))
	}

	if response.Entries[0].Username != "bob" || response.Entries[0].Score != 20 || response.Entries[0].Rank != 1 {
		t.Errorf("Expected 'bob' with score 20 as rank 1, got '%s' with score %d", response.Entries[0].Username, response.Entries[0].Score)
	}
	if response.Entries[1].Username != "carol" || response.Entries[1].Score != 15 || response.Entries[1].Rank != 2 {
		t.Errorf("Expected 'carol' with score 15 as rank 2, got '%s' with score %d", response.Entries[1].Username, response.Entries[1].Score)
	}
	if response.Entries[2].Username != "alice" || response.Entries[2].Score != 10 || response.Entries[2].Rank != 3 {
		t.Errorf("Expected 'alice' with score 10 as rank 3, got '%s' with score %d", response.Entries[2].Username, response.Entries[2].Score)
	}
}
