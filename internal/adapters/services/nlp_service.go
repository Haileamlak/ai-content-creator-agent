package services

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type NLPService struct {
	Client *genai.Client
}

func NewNLPService(apiKey string) *NLPService {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}

	return &NLPService{Client: client}
}

func (service *NLPService) GenerateText(prompt string) (string, error) {
	model := service.Client.GenerativeModel("gemini-1.5-flash")
	model.ResponseMIMEType = "application/json"
	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))

	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	generatedText := formatResponse(resp)
	return generatedText, nil
}

func (service *NLPService) SummarizeText(text string) (string, error) {
	model := service.Client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(context.Background(), genai.Text("Summarize the following text: "+text))

	if err != nil {
		return "", fmt.Errorf("failed to summarize text: %v", err)
	}

	summary := formatResponse(resp)
	return summary, nil
}

func (service *NLPService) AnalyzeText(text string) (string, error) {
	model := service.Client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(context.Background(), genai.Text("Analyze the following text: "+text))

	if err != nil {
		return "", fmt.Errorf("failed to analyze text: %v", err)
	}

	analysis := formatResponse(resp)
	return analysis, nil
}

func formatResponse(resp *genai.GenerateContentResponse) string {
	res := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				res += fmt.Sprint(part) + "\n"
			}
		}
	}

	return res
}
