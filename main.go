package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	API_URL     = "https://api.openai.com/v1/completions"
	API_KEY     = "YOUR_API_KEY"
	MODEL       = "text-davinci-001"
	MAX_TOKENS  = 200
	TEMPERATURE = 0
)

type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Text          string        `json:"text"`
	Index         int64         `json:"index"`
	Logprobs      interface{}   `json:"logprobs"`
	FinishDetails FinishDetails `json:"finish_details"`
}

type FinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

func main() {
	var prompt string

	if len(os.Args) != 1 {
		prompt = strings.Join(os.Args[1:], " ")
		if prompt == "" {
			fmt.Println("No prompt given")
			os.Exit(0)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\033[32mAsk Luna: \033[0m")
		prompt, _ = reader.ReadString('\n')
		prompt = strings.TrimSpace(prompt)
		if prompt == "" {
			fmt.Println("No prompt given")
			os.Exit(0)
		}
		if strings.ToLower(prompt) == "exit" {
			os.Exit(0)
		}
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", API_URL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+API_KEY)

	requestBody := map[string]interface{}{
		"model":       MODEL,
		"prompt":      prompt,
		"max_tokens":  MAX_TOKENS,
		"stream":      true,
		"temperature": TEMPERATURE,
	}

	jsonStr, _ := json.Marshal(requestBody)
	req.Body = ioutil.NopCloser(strings.NewReader(string(jsonStr)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.Contains(line, "data: ") && !strings.Contains(line, "[DONE]") {
			abc := strings.Replace(line, "data: ", "", 1)
			if strings.Contains(abc, "<|im_end|>") {
				abc = strings.Replace(abc, "<|im_end|>", "", 1)
			}
			var response Response
			json.Unmarshal([]byte(abc), &response)
			fmt.Print(response.Choices[0].Text)
		}
	}
}
