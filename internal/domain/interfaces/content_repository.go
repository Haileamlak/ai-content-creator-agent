package interfaces

import (
    "ai-content-creator-agent/internal/domain/entities"
)

// ContentRepository defines the interface for content-related operations.
type ContentRepository interface {
    Save(content *entities.Content) error
    Update(content *entities.Content) error
    Get(companyId string, contentID string) (*entities.Content, error)
	GetPopular(companyID string) (*entities.Content, error)
	GetAll(companyID string, limit int) ([]entities.Content, error)
	Delete(contentID string) error
}