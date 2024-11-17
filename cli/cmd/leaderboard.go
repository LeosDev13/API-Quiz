package cmd

import (
	"fmt"
	"log"
	"quiz-app/shared/client"

	"github.com/spf13/cobra"
)

var leaderboardCmd = &cobra.Command{
	Use:   "leaderboard",
	Short: "View your ranking compared to others",
	Run: func(cmd *cobra.Command, args []string) {
		viewLeaderboard()
	},
}

func init() {
	rootCmd.AddCommand(leaderboardCmd)
}

func viewLeaderboard() {
	apiClient := client.NewAPIClient("http://localhost:3000")

	leaderboard, err := apiClient.FetchLeaderboard()
	if err != nil {
		log.Fatalf("Error fetching leaderboard: %v", err)
	}

	fmt.Printf("\nLeaderboard:\n")
	for _, entry := range leaderboard.Entries {
		fmt.Printf("%s: %d correct answers\n", entry.Username, entry.Score)
	}
}
