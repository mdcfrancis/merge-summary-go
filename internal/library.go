package internal

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
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

type Config struct {
	Qualitative bool
	GptAuth     string
}

func (cfg *Config) getFromGPT(prompt string) (string, error) {
	client := openai.NewClient(cfg.GptAuth)
	log.Println("getting answer from GPT :", len(prompt), "characters")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4Turbo1106, //openai.GPT4, //3Dot5Turbo,
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
func (cfg *Config) SummarizeSummary(summary string) (string, error) {
	prompt := []string{
		"summarize the following summary in a neutral tone,",
		"format the output in markdown",
		"start each section with a header which includes the file name and the type of change",
	}
	if cfg.Qualitative {
		prompt = append(prompt, "add an overall quality summary at the end")
	}
	prompt = append(prompt, summary)
	promptString := strings.Join(prompt, "\n")
	return cfg.getFromGPT(promptString)
}

func (cfg *Config) ChunkToSummary(chunk Chunk) (string, error) {
	prompt := []string{
		"summarize the following diffs in a neutral tone, do not refer to the author,",
		"include the actual file names where appropriate",
		"do not include the actual code changes unless needed for clarity",
		"format the output in markdown",
		"start each section with a header which includes the file name and the type of change",
	}
	if cfg.Qualitative {
		prompt = append(prompt, "Add a short quality analysis of the changes as a separate section")
	}
	prompt = append(prompt, chunk.Content)
	promptString := strings.Join(prompt, "\n")

	return cfg.getFromGPT(promptString)
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

func (cfg *Config) GetDiff(owner string, repo string, pr string) ([]Chunk, error) {
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

func (cfg *Config) GetPRDetail(owner string, repo string, pr string) (PRDetail, error) {
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