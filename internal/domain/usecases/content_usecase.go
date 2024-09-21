package usecases

import (
	"ai-content-creator-agent/internal/adapters/services"
	"ai-content-creator-agent/internal/domain/entities"
	"ai-content-creator-agent/internal/domain/interfaces"
	"ai-content-creator-agent/pkg/utils"
	"encoding/json"
	"log"
	"strings"
	"time"
)

// ContentUsecase is a struct that defines the usecase for content-related operations.
type contentUsecase struct {
	contentRepo interfaces.ContentRepository
	researchUsecase interfaces.ResearchUsecase
	nlpService *services.NLPService
}

// NewContentUsecase creates a new instance of ContentUsecase.
func NewContentUsecase(contentRepo interfaces.ContentRepository, researchUsecase interfaces.ResearchUsecase, nlpService *services.NLPService ) interfaces.ContentUsecase {
	return &contentUsecase{
		contentRepo: contentRepo,
		researchUsecase: researchUsecase,
		nlpService: nlpService,
	}
}

// SaveContent saves the content to the repository.
func (cu *contentUsecase) SaveContent(content *entities.Content) error {
	return cu.contentRepo.Save(content)
}

// UpdateContent updates the content in the repository.
func (cu *contentUsecase) UpdateContent(content *entities.Content) error {
	return cu.contentRepo.Update(content)
}

// GetContent retrieves the content from the repository.
func (cu *contentUsecase) GetContent(companyID string, contentID string) (*entities.Content, error) {
	return cu.contentRepo.Get(companyID, contentID)
}

// GetAllContent retrieves all content from the repository.
func (cu *contentUsecase) GetAllContent(companyID string, limit int) ([]entities.Content, error) {
	return cu.contentRepo.GetAll(companyID, limit)
}

// DeleteContent deletes the content from the repository.
func (cu *contentUsecase) DeleteContent(contentID string) error {
	return cu.contentRepo.Delete(contentID)
}

// GenerateContent generates content for the given company.
func (cu *contentUsecase) GenerateContent(company *entities.Company) (*entities.Content, error) {
	popularContent, err := cu.contentRepo.GetPopular(company.ID)
	if err != nil {
		return nil, err
	}

	research, err := cu.researchUsecase.ResearchNews(company.Industry)
	if err != nil {
		return nil, err
	}

	prompt := "Generate a social media post for the company " +
		"Company Name: " + company.CompanyName + "\n" +
		"Industry: " + company.Industry + "\n" +
		"Mission Statement: " + company.MissionStatement + "\n" +
		"Brand Voice: " + company.BrandVoice + "\n" +
		"Target Audience: " + company.TargetAudience + "\n" +
		"Content Goals: " + strings.Join(company.ContentGoals, ", ") + "\n" +
		"Key Messages: " + strings.Join(company.KeyMessages, ", ") + "\n\n" +
		"Most liked content(But should be unique and different from it): " + popularContent.Message + "\n" + 
		"News Research Summary About The Industry: " + research + "\n" +
		"Make sure the post is ready for immediate publication without any placeholders or markdown formatting. " +
		"The response should contain exactly introduction, body, and hashtagas. in json format."

	generatedContent, err := cu.nlpService.GenerateText(prompt)
	if err != nil {
		return nil, err
	}
	log.Println(generatedContent)

	var contentBody struct {
		Intro string `json:"introduction"`
		Body  string `json:"body"`
		Hashtags string `json:"hashtags"`
	}

	err = json.Unmarshal([]byte(generatedContent), &contentBody)
	if err != nil {
		return nil, err
	}

	id, err := utils.GenerateRandomID(16)

	if err != nil {
		return nil, err
	}

	content := &entities.Content{
		ID:          id,
		CompanyID:   company.ID,
		Message:     contentBody.Intro + "\n\n" + contentBody.Body + "\n\n" + contentBody.Hashtags,
		ContentType: "Social Media Post",
		CreatedAt:   time.Now(),
		Platform:    "Facebook",
		Status:      "Draft",
	}

	return content, nil
}