/*
Copyright © 2024 Michael Francis michael@francis-web.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "merge-summary-go",
	Short: "A set of tools to summarise PRs on GitHub",
	Long:  `A set of tools to summarise PRs on GitHub.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	repoOwner string
	repoName  string
	gptAuth   string
)

func init() {
	rootCmd.Flags().StringVarP(&repoOwner, "owner", "o", "mdcfrancis", "The owner of the repository")
	rootCmd.Flags().StringVarP(&repoName, "repo", "r", "merge-summary-go", "The name of the repository")
	rootCmd.Flags().StringVarP(&gptAuth, "gpt-auth", "a", os.Getenv("GPT_AUTH"), "The GPT API key (not recommended, use environment")
}
