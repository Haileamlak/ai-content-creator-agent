package usecases

import (
	"ai-content-creator-agent/internal/adapters/services"
	"ai-content-creator-agent/internal/domain/interfaces"
	"ai-content-creator-agent/internal/infrastructure/api"
	"log"
	"strings"
)

// ResearchUsecase defines the interface for research-related operations.
type researchUsecase struct {
	nlpService *services.NLPService
	newsAPI *api.NewsAPI
}

// NewResearchUsecase creates a new instance of the research usecase.
func NewResearchUsecase(newsAPI *api.NewsAPI, nlpService *services.NLPService) interfaces.ResearchUsecase {
	return &researchUsecase{
		newsAPI: newsAPI,
		nlpService: nlpService,
	}
}

func (ru *researchUsecase) ResearchNews(industry string) (string, error) {
	// Get the latest news articles
	news, err := ru.newsAPI.SearchNews(industry, 5)
	if err != nil {
		return "", err
	}

	log.Println("News articles retrieved successfully: ", news)
	// Extract relevant information from the news articles using NLP service 
	summary, err := ru.nlpService.SummarizeText(strings.Join(news, "\n"))
	if err != nil {
		return "", err
	}
	log.Println("News articles summarized successfully: ", summary)
	return summary, nil
}