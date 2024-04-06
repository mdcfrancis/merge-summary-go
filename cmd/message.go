/*
Copyright © 2024 Michael Francis michael@francis-web.com
*/
package cmd

import (
	"com.github/mdcfrancis/merge-summary-go/internal"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message <pull-request>",
	Short: "Gets the diffs from a PR and summarises them into a message",
	Long:  `This command will get the diffs from a PR and summarise them into a message.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pullRequest := args[0]
		cfg := internal.Config{
			Qualitative: qualitative,
			GptAuth:     gptAuth,
		}

		detail, err := cfg.GetPRDetail(repoOwner, repoName, pullRequest)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("PR Detail:")
		fmt.Println("Number:", detail.Number)
		fmt.Println("State:", detail.State)
		fmt.Println("Title:", detail.Title)
		fmt.Println("URL:", detail.URL)
		fmt.Println("Diff URL:", detail.DiffURL)

		chunks, err := cfg.GetDiff(repoOwner, repoName, pullRequest)
		if err != nil {
			log.Fatal(err)
		}

		summaries := []string{}
		for _, chunk := range chunks {
			summary, err := cfg.ChunkToSummary(chunk)
			if err != nil {
				log.Fatal(err)
			}
			summaries = append(summaries, summary)
		}

		grouped := strings.Join(summaries, "\n")
		summary, err := cfg.SummarizeSummary(grouped)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(summary)
	},
}

var qualitative bool

func init() {
	rootCmd.AddCommand(messageCmd)
	messageCmd.Flags().BoolVarP(&qualitative, "qualitative", "q", false, "Use qualitative summarization")
}
