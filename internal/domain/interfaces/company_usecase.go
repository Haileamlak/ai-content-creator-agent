package interfaces

import (
	"ai-content-creator-agent/internal/domain/entities"
)

type CompanyUsecase interface {
	RegisterCompany(company *entities.Company) error
	UpdateCompany(company *entities.Company) error
	GetCompany(companyID string) (*entities.Company, error)
	DeleteCompany(companyID string) error
}
