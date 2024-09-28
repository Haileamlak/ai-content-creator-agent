package usecases

import (
	"ai-content-creator-agent/internal/adapters/services"
	"ai-content-creator-agent/internal/domain/entities"
	"ai-content-creator-agent/internal/domain/interfaces"
)

type companyUsecase struct {
	CompanyRepo interfaces.CompanyRepository
	Scheduler   *services.ScheduleService
}

func NewCompanyUsecase(cr interfaces.CompanyRepository, scheduler *services.ScheduleService) interfaces.CompanyUsecase {
	return &companyUsecase{
		CompanyRepo: cr,
		Scheduler:   scheduler,
	}
}

func (uc *companyUsecase) RegisterCompany(company *entities.Company) error {
	// create a new company
	err := uc.CompanyRepo.Create(company)
	if err != nil {
		return err
	}

	// schedule a job to post content
	uri := "<your-backend-server-url>" + "/post-content/" + company.ID
	name := "post-content-" + company.ID
	err = uc.Scheduler.ScheduleJob(uri, name)
	if err != nil {
		return err
	}

	return nil
}

func (uc *companyUsecase) UpdateCompany(company *entities.Company) error {
	return uc.CompanyRepo.Update(company)
}

func (uc *companyUsecase) GetCompany(companyID string) (*entities.Company, error) {
	return uc.CompanyRepo.Get(companyID)
}

func (uc *companyUsecase) DeleteCompany(companyID string) error {
	return uc.CompanyRepo.Delete(companyID)
}
