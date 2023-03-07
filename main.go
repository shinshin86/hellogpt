package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const URL string = "https://api.openai.com/v1/chat/completions"

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type chatResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
}

func chat(messages []message, apiKey string) []message {
	// DEBUG
	// fmt.Println(messages)

	reqBody := chatRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}
	reqJSON, err := json.Marshal(reqBody)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqJSON))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var resBody chatResponse
	json.NewDecoder(res.Body).Decode(&resBody)

	message := resBody.Choices[0].Message

	fmt.Println("ChatGPT: " + message.Content)

	return append(messages, message)
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API key must be defined as an environment variable.")
	}

	messages := []message{}
	isFirstTime := true
	reader := bufio.NewReader(os.Stdin)

	for {
		if isFirstTime {
			fmt.Println("=== What type of chatbot do you want? Please enter the type of chatbot you are looking for. ===")
			text, err := reader.ReadString('\n')

			if err != nil {
				log.Fatalln(err)
			}

			text = strings.TrimSpace(text)
			msg := message{Role: "system", Content: text}
			messages = append(messages, msg)
			isFirstTime = false
			fmt.Println("=== OK! Let's start the conversation. ===")
		} else {
			text, err := reader.ReadString('\n')

			if err != nil {
				log.Fatalln(err)
			}

			text = strings.TrimSpace(text)

			if text == "bye" {
				fmt.Println("bye!")
				os.Exit(0)
			}

			msg := message{Role: "user", Content: text}
			messages = append(messages, msg)
			messages = chat(messages, apiKey)
		}
	}
}
