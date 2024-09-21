package repositories

import (
	"ai-content-creator-agent/internal/domain/entities"
	"ai-content-creator-agent/internal/domain/interfaces"
	"context"

	"cloud.google.com/go/firestore"
)

type CompanyRepository struct {
	Client *firestore.Client
	Ctx    context.Context
}

// NewCompanyRepository creates a new CompanyRepository instance.
func NewCompanyRepository(ctx context.Context, projectID string) interfaces.CompanyRepository {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	
	return &CompanyRepository{
		Client: client,
		Ctx:    ctx,
	}
}

// SaveCompany stores a new Company entity in Firestore.
func (repo *CompanyRepository) Create(Company *entities.Company) error {
	_, err := repo.Client.Collection("Companys").Doc(Company.ID).Set(repo.Ctx, Company)
	return err
}

// UpdateCompany updates an existing Company entity in Firestore.
func (repo *CompanyRepository) Update(Company *entities.Company) error {
	_, err := repo.Client.Collection("Companys").Doc(Company.ID).Set(repo.Ctx, Company)
	return err
}

// FindCompanyByID retrieves a Company entity from Firestore by its ID.
func (repo *CompanyRepository) Get(CompanyID string) (*entities.Company, error) {
	doc, err := repo.Client.Collection("Companys").Doc(CompanyID).Get(repo.Ctx)
	if err != nil {
		return nil, err
	}

	var Company entities.Company
	err = doc.DataTo(&Company)
	if err != nil {
		return nil, err
	}

	return &Company, nil
}

// DeleteCompany removes a Company entity from Firestore by its ID.
func (repo *CompanyRepository) Delete(CompanyID string) error {
	_, err := repo.Client.Collection("Companys").Doc(CompanyID).Delete(repo.Ctx)
	return err
}
