package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-cli",
	Short: "CLI for interacting with the quiz",
	Long:  `A command-line interface to take quizzes and view results.`,
}

func Execute() error {
	return rootCmd.Execute()
}
