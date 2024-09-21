package usecases

import (
	"ai-content-creator-agent/internal/domain/interfaces"
	"ai-content-creator-agent/internal/infrastructure/api"
)

type SocialMediaMgtUsecase struct {
	contentUsecase interfaces.ContentUsecase
	companyUsecase interfaces.CompanyUsecase
	facebookAPI    api.FacebookAPIService
}

func NewSocialMediaMgtUsecase(contentUsecase interfaces.ContentUsecase, companyUsecase interfaces.CompanyUsecase) interfaces.SocialMediaMgtUsecase {
	return &SocialMediaMgtUsecase{
		contentUsecase: contentUsecase,
		companyUsecase: companyUsecase,
	}
}

func (uc *SocialMediaMgtUsecase) PostContent(companyID string) error {
	company, err := uc.companyUsecase.GetCompany(companyID)
	if err != nil {
		return err
	}

	content, err := uc.contentUsecase.GenerateContent(company)
	if err != nil {
		return err
	}

	// post_id, err := uc.facebookAPI.CreatePost(content.Message, company.FacebookPageID, company.FacebookAccessToken)
	// if err != nil {
	// 	return err
	// }

	content.PostID = "post_id"
	err = uc.contentUsecase.SaveContent(content)
	if err != nil {
		return err
	}

	return nil
}

// reply to comments on every post in the company's page in the last 24 hours
func (uc *SocialMediaMgtUsecase) ReplyToComments(companyID string) error {
	company, err := uc.companyUsecase.GetCompany(companyID)
	if err != nil {
		return err
	}

	comments, err := uc.facebookAPI.GetComments(company.FacebookPageID, company.FacebookAccessToken)
	if err != nil {
		return err
	}

	for _, comment := range comments {
		err = uc.facebookAPI.ReplyToComment(comment["id"].(string), "Thank you for your comment!", company.FacebookAccessToken)
		if err != nil {
			return err
		}
	}

	return nil
}