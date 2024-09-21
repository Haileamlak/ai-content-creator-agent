package interfaces

import (
	"ai-content-creator-agent/internal/domain/entities"
)

type ContentUsecase interface {
	SaveContent(content *entities.Content) error
	UpdateContent(content *entities.Content) error
	GetContent(companyID string, contentID string) (*entities.Content, error)
	GetAllContent(companyID string, limit int) ([]entities.Content, error)
	DeleteContent(contentID string) error
	GenerateContent(company *entities.Company) (*entities.Content, error)
}	