package services

import (
	"context"
	"fmt"

	language "cloud.google.com/go/language/apiv1"
	languagepb "cloud.google.com/go/language/apiv1/languagepb"
	"google.golang.org/api/option"
)

// SentimentAnalysisService provides sentiment analysis capabilities for content.
type SentimentAnalysisService struct {
	Client *language.Client
}

func NewSentimentAnalysisService(apiKey string) *SentimentAnalysisService {
	ctx := context.Background()
	c, err := language.NewRESTClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}


	return &SentimentAnalysisService{Client: c,}
}

// AnalyzeSentiment analyzes the sentiment of a given text.
func (service *SentimentAnalysisService) AnalyzeSentiment(text string) (string, error) {
	req := &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	}

	resp, err := service.Client.AnalyzeSentiment(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("failed to analyze sentiment: %v", err)
	}

	if resp.DocumentSentiment.Score < 0 {
		return "Negative", nil
	} else if resp.DocumentSentiment.Score == 0 {
		return "Neutral", nil
	} else {
		return "Positive", nil
	}
}