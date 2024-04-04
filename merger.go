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

// const basePrompt = "summarize the following in a concise manner make it sound like you are the author:\n\n%s\n"
const basePrompt = "summarize the following in a neutral tone, do not refer to the author, include the actual file names, " +
	"use markdown to make the comment clear:\n\n%s\n"

func chunkToSummary(chunk Chunk, gptAuth string) (string, error) {
	prompt := fmt.Sprintf(basePrompt, chunk.Content)

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
	fmt.Println("Body:", detail.Body)
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

	chunk := strings.Join(summaries, "\n")
	summary, err := chunkToSummary(Chunk{Content: chunk}, *gptAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(summary)
}
