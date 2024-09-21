package interfaces

import "ai-content-creator-agent/internal/domain/entities"

type CompanyRepository interface {
    Create(company *entities.Company) error
    Update(company *entities.Company) error
    Get(companyID string) (*entities.Company, error)
    Delete(companyID string) error
}