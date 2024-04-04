package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type PRDetail struct {
	Number            int    `json:"number"`
	State             string `json:"state"`
	Title             string `json:"title"`
	Body              string `json:"body"`
	URL               string `json:"html_url"`
	DiffURL           string `json:"diff_url"`
	PatchURL          string `json:"patch_url"`
	CommentsURL       string `json:"comments_url"`
	ReviewCommentsURL string `json:"review_comments_url"`
}

type Chunk struct {
	Content string
}

func getFromGPT(prompt string, gptAuth string) (string, error) {
	client := openai.NewClient(gptAuth)
	log.Println("getting answer from GPT :", len(prompt), "characters")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4, //openai.GPT4, //3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v\n", err)
	}
	content := resp.Choices[0].Message.Content
	//fmt.Println(content)
	return content, nil
}
func summarizeSummary(summary string, gptAuth string) (string, error) {
	prompt := []string{
		"summarize the following summary in a neutral tone,",
		"format the output in markdown",
		"start each section with a header which includes the file name and the type of change",
		"add an overall quality summary at the end",
		summary,
	}
	promptString := strings.Join(prompt, "\n")
	return getFromGPT(promptString, gptAuth)
}

func chunkToSummary(chunk Chunk, gptAuth string) (string, error) {
	prompt := []string{
		"summarize the following diffs in a neutral tone, do not refer to the author,",
		"include the actual file names where appropriate",
		"format the output in markdown",
		"start each section with a header which includes the file name and the type of change",
	}
	if *qualitative {
		prompt = append(prompt, "Add a short quality analysis of the changes as a separate section")
	}
	prompt = append(prompt, chunk.Content)
	promptString := strings.Join(prompt, "\n")

	return getFromGPT(promptString, gptAuth)
}

func splitDiff(closer io.ReadCloser) ([]Chunk, error) {
	var chunks []Chunk
	var currentChunk Chunk
	scanner := bufio.NewScanner(closer)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "diff --git") {
			// New chunk, save the old one if it's not empty
			if currentChunk.Content != "" {
				chunks = append(chunks, currentChunk)
			}

			// Start a new chunk
			currentChunk = Chunk{Content: line + "\n"}
		} else if currentChunk.Content != "" {
			// Append line to current chunk
			currentChunk.Content += line + "\n"
		}
	}
	// Append the last chunk
	if currentChunk.Content != "" {
		chunks = append(chunks, currentChunk)
	}

	// Handle potential errors from scanner
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Check your chunks
	log.Println("Number of diffs to analyse :", len(chunks))

	return chunks, nil
}

func getDiff(owner string, repo string, pr string) ([]Chunk, error) {
	mediaTypeDiff := "application/vnd.github.v3.diff"
	// set a media type for the diff
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls/"+pr, nil)
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Accept", mediaType)
	req.Header.Set("Accept", mediaTypeDiff)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return splitDiff(resp.Body)
}

func getPRDetail(owner string, repo string, pr string) (PRDetail, error) {
	mediaType := "application/vnd.github.v3+json"
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls/"+pr, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", mediaType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var prDetail PRDetail
	err = json.Unmarshal(body, &prDetail)
	if err != nil {
		log.Fatal(err)
	}
	if prDetail.Number == 0 {
		return prDetail, fmt.Errorf("PR not found")
	}
	return prDetail, nil
}

// add go standard command line args
// using the flags library
// https://golang.org/pkg/flag/
var repoOwner = flag.String("owner", "aws", "The owner of the repository")
var repoName = flag.String("repo", "aws-lambda-go", "The name of the repository")
var prNumber = flag.String("pr", "544", "The pull request number")
var gptAuth = flag.String("gpt", os.Getenv("GPT_AUTH"), "The GPT API key (not recommended, use environment")
var qualitative = flag.Bool("qualative", false, "Use qualitative summarization")

func main() {
	flag.Parse()
	log.Println("owner:", *repoOwner)
	log.Println("repo:", *repoName)
	log.Println("pr:", *prNumber)

	detail, err := getPRDetail(*repoOwner, *repoName, *prNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PR Detail:")
	fmt.Println("Number:", detail.Number)
	fmt.Println("State:", detail.State)
	fmt.Println("Title:", detail.Title)
	//fmt.Println("Body:", detail.Body)
	fmt.Println("URL:", detail.URL)
	fmt.Println("Diff URL:", detail.DiffURL)

	chunks, err := getDiff(*repoOwner, *repoName, *prNumber)
	if err != nil {
		log.Fatal(err)
	}

	summaries := []string{}
	for _, chunk := range chunks {
		summary, err := chunkToSummary(chunk, *gptAuth)
		if err != nil {
			log.Fatal(err)
		}
		summaries = append(summaries, summary)
	}

	grouped := strings.Join(summaries, "\n")
	summary, err := summarizeSummary(grouped, *gptAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(summary)
}
