package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	apiKey, err := getAPIKey()
	if err != nil {
		printError(err)
		printError(errors.New("hint: you can set OPENAI_API_KEY environment " +
			"variable or create askai_config.yaml file with your API key"))
		return
	}

	c := openai.NewClient(apiKey)
	ctx := context.Background()

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			printUsage()
			return
		}
		prompt := ""
		for _, arg := range os.Args[1:] {
			prompt += arg + " "
		}
		simpleMode(c, ctx, prompt)
		return
	}

	fmt.Println("Welcome to askai! Type 'help' to see available commands.")
	history := []openai.ChatCompletionMessage{}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			printError(err)
			continue
		}

		text = strings.TrimSpace(text)

		switch text {
		case "exit":
			return
		case "help":
			printUsage()
			continue
		case "clear":
			history = []openai.ChatCompletionMessage{}
			fmt.Println("Chat history cleared.")
			continue
		default:
			// handle other input cases
		}

		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 500,
			Messages:  history,
			Stream:    true,
		}
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			printError(err)
			continue
		}
		historyText := ""
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				stream.Close()
				break
			}

			if err != nil {
				printError(err)
				stream.Close()
				break
			}
			delta := response.Choices[0].Delta.Content
			fmt.Print(delta)
			historyText += delta
		}
		fmt.Println()
		history = append(history, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: historyText,
		})
	}
}

func printUsage() {
	fmt.Println("Available commands:")
	fmt.Println("  help - show this message")
	fmt.Println("  exit - exit the program")
	fmt.Println("  clear - clear chat history")
}

func getAPIKey() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey != "" {
		return apiKey, nil
	}

	file, err := os.Open("askai_config.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "api_key:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "api_key:")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	return "", errors.New("API key not found")
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "\x1b[31mError: %v\x1b[0m\n", err)
}

// simpleMode is a simple mode of askai that takes a prompt and prints the result to stdout.
func simpleMode(c *openai.Client, ctx context.Context, prompt string) {

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		printError(fmt.Errorf("chatCompletionStream error: %v", err))
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			printError(fmt.Errorf("\nstream error: %v", err))
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
