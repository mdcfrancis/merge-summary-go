/*
Copyright © 2024 Michael Francis michael@francis-web.com
*/
package cmd

import (
	"com.github/mdcfrancis/merge-summary-go/internal"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

const gitHubURL = "https://github.com"

// docCmd represents the doc command
var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Given a owner and repo name generate a summary document of the repo ",
	Long:  `Given a owner and repo name generate a summary document of the repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo := fmt.Sprintf("%s/%s/%s", gitHubURL, repoOwner, repoName)
		// Clone the given repository to memory
		log.Println("Cloning repository", repo)
		r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: repo,
		})
		if err != nil {
			log.Fatal(err)
		}
		ref, err := r.Head()
		if err != nil {
			log.Fatal(err)
		}
		commit, err := r.CommitObject(ref.Hash())
		if err != nil {
			log.Fatal(err)
		}

		tree, err := commit.Tree()
		if err != nil {
			log.Fatal(err)
		}
		cfg := internal.Config{
			GptAuth: gptAuth,
		}

		excludedFiles := []string{"go.sum", "go.mod", "LICENSE", ".md"}

		ft := internal.FileTools{
			Directory: outputDirectory,
		}
		// Check that the output directory exists if it doesn't create it
		err = ft.CheckOrCreateDirectory()
		summaries := []string{}
		err = tree.Files().ForEach(func(f *object.File) error {
			for _, excluded := range excludedFiles {
				if strings.HasSuffix(f.Name, excluded) {
					return nil
				}
			}
			// Exclude files in .directories
			if f.Name[0] == '.' {
				return nil
			}
			log.Println("Processing file", f.Name)
			parts := []string{f.Name}
			lines, err := f.Lines()
			if err != nil {
				return err
			}
			parts = append(parts, lines...)
			allLines := strings.Join(parts, "\n")

			summary, err := cfg.FileToSummary(internal.Chunk{Content: allLines})
			if err != nil {
				return err
			}
			// create a filename, replacing the path separators with dots and lower casing, then adding .md
			fileName := strings.ToLower(strings.ReplaceAll(f.Name, string(os.PathSeparator), ".")) + ".md"
			prefixed := []string{"Doc File:" + fileName, "File name:" + f.Name, "Summary:" + summary}
			summaries = append(summaries, strings.Join(prefixed, "\n"))
			err = ft.WriteFile(fileName, summary)
			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
		log.Println("Creating README.md")
		summary, err := cfg.IndexPage(repo, outputDirectory, summaries)
		if err != nil {
			log.Fatal(err)
		}
		err = ft.WriteFile("README.md", summary)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var outputDirectory string

func init() {
	rootCmd.AddCommand(docCmd)
	docCmd.Flags().StringVarP(&outputDirectory, "output", "o", "", "The directory to output the summary")
	docCmd.MarkFlagRequired("output")
}
