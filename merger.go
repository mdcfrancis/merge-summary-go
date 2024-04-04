package main

import (
	"bufio"
	"context"
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
const basePrompt = "summarize the following and make it sound like you are the author, include the actual file names:\n\n%s\n"

func chunkToSummary(chunk Chunk) (string, error) {
	prompt := fmt.Sprintf(basePrompt, chunk.Content)
	gptAuth := os.Getenv("GPT_AUTH")
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
	for i, chunk := range chunks {
		fmt.Printf("Chunk %d:\n%s\n", i+1, chunk.Content)
	}

	return chunks, nil
}

func main() {
	repoOwner := "aws"
	repoName := "aws-lambda-go"
	prNumber := "544"

	//mediaType := "application/vnd.github.v3+json"
	mediaTypeDiff := "application/vnd.github.v3.diff"
	// set a media type for the diff
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+repoOwner+"/"+repoName+"/pulls/"+prNumber, nil)
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

	chunks, err := splitDiff(resp.Body)

	//body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	summaries := []string{}
	for _, chunk := range chunks {
		summary, err := chunkToSummary(chunk)
		if err != nil {
			log.Fatal(err)
		}
		summaries = append(summaries, summary)
	}

	chunk := strings.Join(summaries, "\n")
	summary, err := chunkToSummary(Chunk{Content: chunk})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(summary)

	/*
		var prDetail PRDetail
		err = json.Unmarshal(body, &prDetail)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("PR Detail: %+v\n", prDetail)

	*/
}
